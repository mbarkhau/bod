# bod
Save your RAM and use Bytes-On-Disk

## General Idea
Instead of having big datastrcutures in memory, one can have them mapped as
bytes on disk. The idea is to have a concurrent thread safe "bytes manager"
that communicate to open files on disk.