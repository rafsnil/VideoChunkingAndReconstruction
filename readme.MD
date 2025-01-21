
# Video Chunking and Reconstruction

A simple project that splits a video into equal chunks and store it in different directories (simulating nodes in a distributed system).
The purpose of this is to simply demonstrate how a video can be chunked and stored in a distributed system and then reconstructed from the chunks.

**Functionalities:**
- can chunk the video into specified sizes and store in the different nodes.
- can reconstruct the video from the chunks and give the output in the Output directory.


**Notes:**
- This is a very simple procedure to chunk, store and retrieve.
- This is definitely not an efficient approach for large video files.


**Further Improvements:** <br>
To improve efficiency, consider the following optimizations:
- Use streaming to process chunks without loading them entirely into memory.
- Implement parallel processing to handle multiple chunks simultaneously.
- Optimize I/O operations to reduce overhead.
- Implement robust error handling and recovery mechanisms.
- Ensure efficient use of disk space and network resources.