1. Given a graph, find if there is a path between 2 nodes
Algo: Use DFS (easy to implement with recursion) OR BFS(better when finding shortest path, and quicker sometimes)

2. DFS Traversal - use stack
Algo:   
Step 1: Push the root node in the Stack.  
Step 2: Loop until stack is empty. 
Step 3: Peek the node of the stack.  
Step 4: If the node has unvisited child nodes, get the unvisited child node, mark it as traversed and push it on stack.   
Step 5: If the node does not have any unvisited child nodes, pop the node from the stack.

BFS - Use Queue
Algo:
Step 1: Push the root node in the Queue.
Step 2: Loop until the queue is empty.
Step 3: Remove the node from the Queue.
Step 4: If the removed node has unvisited child nodes, mark them as visited and insert the unvisited children in the queue.
