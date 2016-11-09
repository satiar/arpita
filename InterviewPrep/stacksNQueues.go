1. Store 3 stacks in a single array
Algo: Start storing first stack form a[0]
Store second stack from a[len-1] backwards
Store thirs stack starting from the 'mid' and alternating, mid-1, mid+1, mid-2, mid+2 so on

2. Have a "min" mehod for a stack, alongwith regular push and pop
Algo: Maintain a separate stack s.t. for every 'push' in original stack, 
calculate the local min, and push onto the second stack.
So when a pop happens, pop the corresponding Min from second stack.

3. Stack of plates
Algo: Maintain a list of stacks
Push = getLastStack -> if lastStack != null && lastStack.IsnotFull , lastStack.push(value) -> else create new stack 
and push onto it -> add new stack to list of stacks

Pop = same as regular pop except check if stack gets empty after popping this element..if yes, remove that stack form list of stacks

4. Implement queue with 2 stacks.
Algo: At start - Stack 1 has the elements in stack-order. Stack 2 is empty.
If stack 2 is empty, push all elements of stack 1 onto stack2 in reverse order.
When want to get the first element(FIFO), pop from second stack.
When want to push onto queue, just push onto stack as would do normally.
Again, if stack 2 is empty, push all elements of stack1 in reverse order.

type stack Struct {
}
type Queue Struct {
  stackNew stack
  stackOld stack
}

//Add to Queue is equal to push into stackNew

//MoveElements from New to Old i.e copy all elements from new to old, in reverse order, since Queue is FIFO
func moveElems(stackNew, stackOld stack) {
  if stackOld == empty {
    stackOld.push(stackNew.Pop())
  }
}

//Remove from Queue == pop from stackOld

5. Sort a stack using only one other stack. O(N2)
Algo:  Create a new stack ->Pop top element from stack1->place this popped element into a tmp variable -> pop all elements in 
stack2 which are greater than tmp and push onto stack1 -> push tmp onto stack2 ->  repeat unless stack1 is empty
//pseudo code
func sort(s stack) stack {
  n := Stack{}
  for len(s) > 0 {
    tmp := s.Pop()
    //Keep popping elems from second stack if they are greater than tmp
    for len(n) >0 && s.Peek() > tmp {
      s.push(r.pop())
    }
    // All greater elems have been pushed onto stack s, so push tmp onto sorted stack 2
    r.push(tmp)
  }
  //place back elems from stack2 to stack1
