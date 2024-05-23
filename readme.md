# Concurrent Data Structures in Go

This repository contains implementations of various concurrent data structures in Go. These data structures are designed to be safe for concurrent access by multiple goroutines.

### Note:

This repository is a practice project based on the concurrency part of the Operating Systems: Three Easy Pieces (OSTEP) book. The code examples in the book are written in C, while we have implemented them in Go for practice.

## Introduction

Concurrent data structures are essential in concurrent programming to ensure safe access and manipulation of shared data by multiple concurrent processes or threads. Go's built-in concurrency primitives, such as goroutines and channels, make it convenient to build efficient and safe concurrent data structures.

This repository provides implementations of the following concurrent data structures:

- Concurrent Linked List
- Concurrent Counter
- Concurrent Hash Table(TODO)
- Concurrent Queue(TODO)

Each data structure is designed to support concurrent operations efficiently while ensuring thread safety.

## Implemented Data Structures

### 1. Concurrent Linked List

The concurrent linked list allows concurrent insertion and lookup operations Deletion, etc... . It uses fine-grained locking to ensure thread safety.

### 2. Concurrent Counter

The concurrent counter provides a scalable counter implementation that allows multiple threads to increment and retrieve the counter value concurrently. It uses local and global counters to minimize contention.
