# Tree

A **binary tree** is a tree where every node is connected to, at most, two childen.
In a **complete tree**, all the tree's levels are completely filled, except possibly the last level, where all the nodes are pushed to the left. A **complete binary tree** can be stored with an array easily: Please the root node at index 0. Then for any node with index i, place its children at indices 2xi+1 and 2xi+2.
A **heap** is a complete binary tree where every node holds a value that is at least as large as the values in all its children.