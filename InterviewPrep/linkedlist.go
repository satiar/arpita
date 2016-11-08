1. Merge 2 sorted linked lists into one
//FOLLOWING IS NOT GOLANG CODE
struct node* SortedMerge(struct node* a, struct node* b) 
{
  struct node* result = NULL;
 
  /* Base cases */
  if (a == NULL) 
     return(b);
  else if (b==NULL) 
     return(a);
 
  /* Pick either a or b, and recur */
  if (a->data <= b->data) 
  {
     result = a;
     result->next = SortedMerge(a->next, b);
  }
  else
  {
     result = b;
     result->next = SortedMerge(a, b->next);
  }
  return(result);
}

2. Remove duplicates from a linked list O(N)
Algo 1 : for each element of list, check if it exists in map, if yes, (which means it has occured once, 
and this current one is a duplicate) remove it (prev.next = curr.next), else add it to the map.
Algo 2 :(no additional buffer) - for each current element, go delete all subsequent elements with same value
So keep a curr , and a runner which starts with curr , and goes on till 
while runner.next !=null {
  if runner.next.value == curr.value {
    //delete it
    runner.next = runner.next.next 
  } else {
      runner = runner.next
  }
  current = current.next
  
3. Return nth from last element of a linked list.
  Algo: Start 2 pointers, and advance he second by "n". Now when p2 hits the end, p1 is at n distance from end. so p1 is the 
  nth last element
// pseudo code
LinkedListNode nthToLast(LinkedListNode head, int n) { 
  if (head == null || n < 1) { 
   return null; 
  }
  LinkedListNode p1 = head; 
  LinkedListNode p2 = head; 
  for (int j = 0; j < n - 1; ++j) { // skip n-1 steps ahead 
    if (p2 == null) { 
      return null; // not found since list size < n 
    } 
    p2 = p2.next; 
  } 
  while (p2.next != null) { 
    p1 = p1.next; 
    p2 = p2.next; 
  } 
  return p1; 
}

4. Delete a node with access to only that node
Algo : Since there is no head etc available, assign the next node's pointer and data to this node, hereby indirectly deleting
this node - IMPOSSIBLE IS NODE IS THE LAST NODE
  next := thisNode.next
  thisNode.Data = next.Data
  thisNode.Next = next.Next
  
  5. Add two numbers represented by a linked list (reversed)   1->2->3 + 6->1->4 = 321+ 416 = 7->3->7
  Algo : Solve recursively. Start at the head. add and recurse sending the carry-over to next call.
  //pseudo code 
  addList(l1, l2 *node, carry int) {
    if l1 == null && l2 == null && carry == 0 {
        return null
    }
    //create a new node
    var result Node{}
    sum := carry
    if l1 != null {
        sum += l1.data
    }
    if l2 != null {
      sum += l2.data
    }
    //get the value to be put in result
    result.Data = sum % 10
    //recurse 
    carry := 0
    if sum > 10 {
        carry = 1
    }
    nextNode := addList(l1.next (SEND NULL IF IT IS NULL), l2.next, carry)
    result.next = nextNode
    
  }
  return result
}
    
