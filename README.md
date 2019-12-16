# gostack
Implementation of Stack using Array and Linked list in Go.

Stack follows Last in First out principal to insert and delete elements. Elements are inserted and deleted from the one end only, called Top. 

Generally folowing operations are performed on stack:-
--- Push -> Insert Element into stack.
--- Pop  -> Delete an element from stack.
--- Peek -> Get given element from stack(This does not delete the element).
--- Top  -> Get the top element of the stack(This operation does not delete the top element but just fetches it).
--- IsEmpty -> Checks whether the stack is empty or not. (StackUnderFlow?)
--- IsFull -> Checks whether the stack is full or not. (StackOverflow?)
--- Display -> Prints the elements of the stack.

Each operation(except Display) in stack takes O(1) time. Display takes O(n) time since whole list need to be traversed in order to print it.
Generally stack is implemented using either array or linked list.
