# [622. Design Circular Queue](https://leetcode.com/problems/design-circular-queue/)

Design your implementation of the circular queue. The circular queue is a linear data structure in which the operations are performed based on FIFO (First In First Out) principle and the last position is connected back to the first position to make a circle. It is also called "Ring Buffer".

One of the benefits of the circular queue is that we can make use of the spaces in front of the queue. In a normal queue, once the queue becomes full, we cannot insert the next element even if there is a space in front of the queue. But using the circular queue, we can use the space to store new values.

Your implementation should support following operations:

- `MyCircularQueue(k)`: Constructor, set the size of the queue to be k.
- `Front`: Get the front item from the queue. If the queue is empty, return -1.
- `Rear`: Get the last item from the queue. If the queue is empty, return -1.
- `enQueue(value)`: Insert an element into the circular queue. Return true if the operation is successful.
- `deQueue()`: Delete an element from the circular queue. Return true if the operation is successful.
- `isEmpty()`: Checks whether the circular queue is empty or not.
- `isFull()`: Checks whether the circular queue is full or not.
 

**Example:**
```
MyCircularQueue circularQueue = new MyCircularQueue(3); // set the size to be 3
circularQueue.enQueue(1);  // return true
circularQueue.enQueue(2);  // return true
circularQueue.enQueue(3);  // return true
circularQueue.enQueue(4);  // return false, the queue is full
circularQueue.Rear();  // return 3
circularQueue.isFull();  // return true
circularQueue.deQueue();  // return true
circularQueue.enQueue(4);  // return true
circularQueue.Rear();  // return 4
``` 
**Note:**

- All values will be in the range of [0, 1000].
- The number of operations will be in the range of [1, 1000].
- Please do not use the built-in Queue library.

## 思路

循环队列，队列保持不变，添加或移除元素时更新首尾索引。

初始化：首尾索引都是 -1，并以此判断队列是否为空。  
添加元素时：如果首索引是 -1 则首索引改为 0；尾索引移到下一个位置的索引，并将这个位置的值修改。  
移除元素时：如果首尾索引相同表明只有一个元素，移除后变成空，所以把首尾索引都改为 -1，否则首索引移到下一个位置的索引（没必要删除元素）。  
计算下一个位置的索引：如果索引 +1 等于队列长度，则索引应该回到 0，否则为索引 +1。  
判断是否空：头索引等于 -1。  
判断是否满：尾索引的下一个位置的索引等于头索引。
