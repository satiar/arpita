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


3. Given a sorted arrya, create a minimalist binary search tree
ALgo: TO be minimalist, height of left should be almost same as height of right..so mid should be root
func createTree(arr []int, start, end int) *Node {
  mid := (start + end )/2
  //New node
  node.data = arr[mid]
  node.left = createTree(arr, start, mid)
  node.right = createTRee(arr, mid+1, end)
  return node
}

4. Given a binary tree, design an algo to create a linked list of all elements at all the depths. i.e if there tree is n-deep, have n 
linked lists

Algo: Pre-order traversal (DFS) of tree while also passing the level info
Create a map of level -> list of nodes at that level.

func createList(root *Node, levelMap map[int][]int, level int) {
  if root == null {
      return
  }
  exists, val := levelMap[level]
  if !exists {
    //create a new list
    //add node to new list
    //add list to Map
  }  else {
    //add node to existing list
  }
  createList(root.Left, levelMap, level+1)
  createList(root.Right, levelMap, level+1)
}

5. Balanced Tree (when difference in height of subtrees <=1). Check if a given tree is so.
func checkHtight(root *Node) int {
  if root == null {
    return -1
  }
  left := checkHeight(root.Left)
  right := checkHeight(root.Right)
  diff := left - right
  if math.Abs(diff) >1 {
    return -1
  } else {
    return max(left, right) + 1
  }
}

6. Check is a binary tree is a Binary search tree
Algo: Start with min, max == 0 . If current node is greater than min and less than max, 
then recurse thru left child setting the max = curr data value
then recurse thru right child with min = curr data value

func checkBST(node * Node, min, max int) bool {
  if node == null {
    return true
  }
  if min != 0 && node.Data <min || (max !=0 && node.Data >max) {
    return false
  }
  if !checkBST(node.Left, min, node.Data) || !checkBST(node.Right, node.Data, max) {
      return false
  }
  return true
}
    
    
    
  
