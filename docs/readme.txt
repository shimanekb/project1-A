# input sample

`input_sample.txt`: This file includes the sample inputs for your key-value store. The first line is the decripter of columns.

`input.txt`: This is the for your testing.

descripter: type,key1,key2,value

columns:
- type: There are four types of operations. `put`, `get`, `scan` and `del`.
- key1: All operation types have content in this column. It is a 16-byte string.
- key2: Only `scan` type has content in this column. It is a 16-byte string.
- value: Only `put` type has content in this column. It is a 16-byte string.

# Note

For this project, you only have to implement put/get/del.

