# code-buddy-ai
## Overview

Code Buddy is an intelligent assistant designed to help developers quickly understand and navigate large codebases. By using AI-powered technology, it provides answers to queries regarding project structure, file contents, and specific functions within the code. It's like having a personal assistant dedicated to your code!

## Features

- Codebase Query Assistant: Ask questions about the code, and get instant answers.
- Searchable Documentation: Search through documentation, files, and functions to understand the logic of your codebase.
- Code Embedding: Converts the codebase into embeddings for easier retrieval and understanding.
- Scalable Architecture: The app uses Kafka and Kubernetes to scale and manage processing heavy tasks such as embedding generation.

## Tech Stack

- Frontend: Built with Next.js.
- Backend API: REST API Developed in Go.
- Backend Consumers: Built using Python and Langchain.
- Communication: gRPC - Connection between consumer and RestAPI to stream responses to client.
- Messaging: Kafka.
- Data Storage: Document Store, Vector stores (e.g., Pinecone, FAISS) for managing embeddings.


## System Architecture:

![Miro Untitled](https://github.com/user-attachments/assets/8f6d5ceb-e3c3-4097-b22a-a8cbf780713e)
