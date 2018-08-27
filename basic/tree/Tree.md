# Tree

A tree consists of *nodes* connected by *branches*. Usually the nodes contain some sort of data, and the branches do not. 

Normally branches are drawn as arrows pointing from the *parent node* to the *child node*. Two nodes that have the same parent are sometimes called *siblings*. Each node in the tree has exactly one parent node, except for a single *root node*, which has no parent.

The children, the children's children, and so on for a node are that node's *descendants*. A node's parent, its parent's parent, and so on up to the root are that node's *ancestors*.

Depending on the type of tree, nodes may have any number of children. The number of children a node has is the node's *degree*. A tree's degree is the maximum degree of its nodes. For example, in a degree 2 tree, which is usually called a *binary tree*, each node can have at most two children.

A node with no children is called a *leaf node* or an *external node*. A node that has at least one child is called an *internal node*.

A node's *level* or *depth* in the tree is the distance from the node to the root. In particular, the root's level is 0. A node's *height* is the length of the longest path from the node downward through the tree to a leaf node. In other words, it's the distance from the node to the bottom of the tree.

A *full tree* is one in which every node has either zero children or as many children as the tree's degree. For example, in a full binary, every node has either zero or two children.

A *complete tree* is one in which every level is completely full, except possibly the bottom level, where all the nodes are pushed as far to the left as possible.

A *perfect tree* is full, and all the leaves are at the same level. In other words, the tree holds every possible node for a tree of its height.

## Binary tree

A **binary tree** is a tree where every node is connected to, at most, two childen. It has the following useful facts.

- The number of branches B in a binary tree containing N nodes is B = N − 1.
- The number of nodes N in a perfect binary tree of height H is N = 2^(H+1) − 1.
- Conversely, if a perfect binary tree contains N nodes, it has a height of log(N + 1) − 1.
- The number of leaf nodes L in a perfect binary tree of height H is L = 2^H. Because the total number of nodes in a perfect binary tree of height H is 2^(H+1) − 1, the number of internal nodes I is I = N − L = (2^(H+1) − 1) − 2^H = (2^(H+1) − 2^H) − 1 = 2^H × (2 − 1) − 1 = 2^H − 1.
- This means that in a perfect binary tree, almost exactly half of the nodes are leaves and almost exactly half are internal nodes. More precisely, I = L − 1.
- The number of missing branches (places where a child could be added) M in a binary tree that contains N nodes is M = N + 1.
- If a binary tree has N0 leaf nodes and N2 nodes with degree 2, N0 = N2 + 1. In other words, there is one more leaf node than nodes with degree 2.

A **complete binary tree** can be stored with an array easily: Please the root node at index 0. Then for any node with index i, place its children at indices 2xi+1 and 2xi+2.

A **heap** is a complete binary tree where every node holds a value that is at least as large as the values in all its children.