package main

/**
有一个包含 20 亿个全是 32 位整数的大文件
在其中找到出现次数最多的数，内存限制为2GB
*/

// 每个整数 4B，maxInt32 > 20亿，如果使用 hash 表，则 key 和 value 都需要 4B
// 20 亿个整数，假设全不相同，hash 表需要 20 亿 * 8B = 160亿B = 16GB 空间
// 可以先将 20 亿个整数分成 20 份，每份约 1 亿个整数，这里需要一个优秀的 hash 函数
// 相同的数会被分到同一份，然后统计每份中出现次数最多的数，最后在这 20 个数中找到出现次数最多的数
