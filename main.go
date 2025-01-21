package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const (
	chunkSize = 2 * 1024 * 1024 // 2 MB
)

func main() {

	nodes := []string{"node1", "node2", "node3"}

	videoPath := "Input/input_video.webm"
	if err := createNodeDirectories(nodes); err != nil {
		fmt.Printf("Error creating node directories: %v\n", err)
		return
	}

	if err := chunkAndDistributeVideo(videoPath, nodes); err != nil {
		fmt.Printf("Error chunking and distributing video: %v\n", err)
		return
	}

	fmt.Println("Video chunking complete!")

	retrievedVideoPath := "Output/input_video.webm"
	if err := retrieveAndReassembleVideo(retrievedVideoPath, nodes); err != nil {
		fmt.Printf("Error retrieving and reassembling video: %v\n", err)
		return
	}

	fmt.Println("Video retrieval complete!")

	//// Delete the Nodes
	//if err := deleteNodes(nodes); err != nil {
	//	fmt.Printf("Error deleting node directories: %v\n", err)
	//	return
	//}
}

func createNodeDirectories(nodes []string) error {
	for _, node := range nodes {
		if err := os.MkdirAll(node, os.ModePerm); err != nil {
			return fmt.Errorf("error creating directory %s: %v", node, err)
		}
	}
	return nil
}

func chunkAndDistributeVideo(videoPath string, nodes []string) error {
	file, err := os.Open(videoPath)
	if err != nil {
		return fmt.Errorf("error opening video file: %v", err)
	}
	defer file.Close()

	chunkIndex := 0
	for {
		// Allocate a buffer of size chunkSize to hold the data read from the file
		buffer := make([]byte, chunkSize)

		// Read up to chunkSize bytes from the file into the buffer
		bytesRead, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			return fmt.Errorf("error reading file: %v", err)
		}

		// If no more bytes are read, break the loop
		if bytesRead == 0 {
			break
		}

		// Select a node in a round-robin fashion
		node := nodes[chunkIndex%len(nodes)]

		// Create the path for the chunk file
		chunkPath := filepath.Join(node, fmt.Sprintf("chunk_%d", chunkIndex))

		// Write the chunk of data to the chunk file
		if err := os.WriteFile(chunkPath, buffer[:bytesRead], os.ModePerm); err != nil {
			return fmt.Errorf("error writing chunk to %s: %v", chunkPath, err)
		}

		// Log the successful storage of the chunk
		fmt.Printf("Chunk %d stored in %s\n", chunkIndex, chunkPath)

		// Increment the chunk index for the next iteration
		chunkIndex++
	}
	return nil
}

func retrieveAndReassembleVideo(retrievedVideoPath string, nodes []string) error {
	retrievedFile, err := os.Create(retrievedVideoPath)
	if err != nil {
		return fmt.Errorf("error creating retrieved video file: %v", err)
	}
	defer retrievedFile.Close()

	chunkIndex := 0
	for {
		node := nodes[chunkIndex%len(nodes)]
		chunkPath := filepath.Join(node, fmt.Sprintf("chunk_%d", chunkIndex))

		chunkData, err := os.ReadFile(chunkPath)
		if err != nil {
			if os.IsNotExist(err) {
				break
			}
			return fmt.Errorf("error reading chunk from %s: %v", chunkPath, err)
		}

		if _, err := retrievedFile.Write(chunkData); err != nil {
			return fmt.Errorf("error writing chunk to retrieved video file: %v", err)
		}

		chunkIndex++
	}
	return nil
}

func deleteNodes(nodes []string) error {
	for _, node := range nodes {
		if err := os.RemoveAll(node); err != nil {
			return fmt.Errorf("error deleting directory %s: %v", node, err)
		}
	}
	return nil
}
