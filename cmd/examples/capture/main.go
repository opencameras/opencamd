package main

import "github.com/opencameras/opencamd/pkg/capture"

/* Valgrind report summary

==13531== LEAK SUMMARY:
==13531==    definitely lost: 0 bytes in 0 blocks
==13531==    indirectly lost: 0 bytes in 0 blocks
==13531==      possibly lost: 1,440 bytes in 5 blocks
==13531==    still reachable: 0 bytes in 0 blocks
==13531==         suppressed: 0 bytes in 0 blocks

*/

func main() {
	capture.ListCameras()
}