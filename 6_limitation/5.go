package main

/**
32 位无符号整数的范围是0〜4294967295，现在有 40 亿个32 位无符号整数， 可以使用最多 1GB 的内存， 找出所有出现了两次的数

可以使用最多 10MB 的内存， 怎么找到这 40 亿个整数的中位数
*/

// 1. 申请一个长度为 4294967296 * 2 的 bitset（1GB），初始都是 0
// 2. 对于数 n，第一次遇到将第 2*n 与 2*n+1 位置设为 01，第二次遇设为 10，之后再遇到均设为 11
// 3. 遍历 bitset，找到为 10 的位置，即为出现了两次的数

// 1. 长度为 2MB 的无符号整型数组占用的空间为 8MB
// 2. 4294967295/2M = 2148, 将 32 位无符号整数分为 2148 个区间
// 3. 遍历文件，统计每个区间的数目，找到中位数所在的区间
// 4. 遍历文件，只关注刚刚找到区间上的数，统计每个数的数目
// 5. 根据前面区间的数目，找到中位数
