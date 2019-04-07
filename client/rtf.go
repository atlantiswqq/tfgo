package main

import (
	"context"
	"flag"
	"github.com/galeone/tfgo/rtf"
	"google.golang.org/grpc"
	"io"
	"log"
)

var (
	clientAddr = flag.String("client_addr", "127.0.0.1:50051", "The client address in the format of host:port")
)

func graph(client rtf.RTFClient) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	stream, err := client.DefineAndCall(ctx)
	if err != nil {
		log.Fatalf("%v.DefineAndCall(_) = _, %v", client, err)
	}

	stmts := []string{
		"model = tf.keras.Sequential([tf.keras.layers.Dense(5, activation=tf.nn.elu), tf.keras.layers.Dense(1)])",
		"y_true = tf.constant([[100.0] * 128])",
		"opt = tf.optimizers.Adam()",
		"for i in tf.range(100):",
		"    with tf.GradientTape() as tape:",
		"        y_pred = model(y_true)",
		"        loss = tf.losses.MSE(y_true, y_pred)",
		"        print(\"loss: \", loss.numpy())",
		"    gradients = tape.gradient(loss, model.trainable_variables)",
		"    opt.apply_gradients(zip(gradients, model.trainable_variables))",
	}

	// Sending all requests
	go func() {

		for _, stmt := range stmts {
			if err := stream.Send(&rtf.RTFStatement{
				Stmt: stmt,
			}); err != nil {
				log.Fatalf("Failed to send message: %v", err)
			}
		}
		// Sent everything, close the stream
		if err := stream.CloseSend(); err != nil {
			log.Fatalf("Unable to CloseSend stream: %v", err)
		} else {
			log.Println("Tensorflow function sent")
		}

	}()

	waitc := make(chan struct{})
	// Receiving all the Responses from Tensorflow
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// Received last response
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive message: %v", err)
			}
			// Received response, print it
			if in.Stdout != "" {
				log.Print(in.Stdout)
			}
		}
	}()

	<-waitc
}

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(*clientAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := rtf.NewRTFClient(conn)
	graph(client)
}
