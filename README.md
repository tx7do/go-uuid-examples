# go-uuid-examples

| 软件包                                                            | 样例                                     | 格式                                                                                 |
|----------------------------------------------------------------|----------------------------------------|------------------------------------------------------------------------------------|
| [ksuid](github.com/segmentio/ksuid)                            | 0pPKHjWprnVxGH7dEsAoXX2YQvU	           | 4 bytes of time (seconds) + 16 random bytes                                        |
| [xid](github.com/rs/xid)                                       | b50vl5e54p1000fo3gh0                   | 4 bytes of time (seconds) + 3 byte machine id + 2 byte process id + 3 bytes random |
| [betterguid](github.com/kjk/betterguid)                        | -N-35bz_wVbhxzTZD2wO                   | 8 bytes of time (milliseconds) + 9 random bytes                                    |
| [sonyflake](github.com/sony/sonyflake)                         | 402329004549341446                     | ~6 bytes of time (10 ms) + 1 byte sequence + 2 bytes machine id                    |
| [snowflake](github.com/bwmarrin/snowflake)                     | 1512052461246971904                    | 1 Bit Unused + 41 Bit Timestamp + 10 Bit NodeID + 12 Bit Sequence ID               |
| [snowflake](github.com/godruoyi/go-snowflake)                  | 1512052461246971904                    | 1 Bit Unused + 41 Bit Timestamp + 10 Bit NodeID + 12 Bit Sequence ID               |
| [ulid](github.com/oklog/ulid)                                  | 01BJMVNPBBZC3E36FJTGVF0C4S	            | 6 bytes of time (milliseconds) + 8 bytes random                                    |
| [sid](github.com/chilts/sid)                                   | 1JADkqpWxPx-4qaWY47~FqI	               | 8 bytes of time (ns) + 8 random bytes                                              |
| [go.uuid](github.com/satori/go.uuid)                           | 5b52d72c-82b3-4f8e-beb5-437a974842c	   | UUIDv4 from RFC 4112 for comparison                                                |
| [google.uuid](https://github.com/google/uuid)                  | dd5f48eb-1722-4e5f-9d56-dcaf0aae1026		 | UUIDv4                                                                             |
| [bomberman.uuid](https://github.com/pborman/uuid)              | 330806cf-684c-43f7-85aa-79939ed3415d		 | UUIDv4                                                                             |
| [MongoDB ObjectID](go.mongodb.org/mongo-driver/bson/primitive) | 624ee0fb37583f00042dc309		             | 4-byte timestamp + 5-byte random value + 3-byte incrementing counter               |

## 参考资料

* [Generating good unique ids in Go](https://blog.kowalczyk.info/article/JyRZ/generating-good-unique-ids-in-go.html)
* [Generate a UUID/GUID in Go (Golang)](https://golangbyexample.com/generate-uuid-guid-golang/)
