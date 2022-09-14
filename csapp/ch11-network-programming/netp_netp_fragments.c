
// _t 后缀 表示这些数据类型是用 typedef 定义的, 而不是新的数据类型
// _in 后缀 是 internet 的缩写

// ip socket address structure
struct sockaddr_in {
  uint16_t      sin_family;       // protocol family, always AF_INET
  uint16_t      sin_port;         // port number in network byte order
  struct        in_addr sin_addr; // ip address  in network byte order
  unsigned char sin_zero[8];      // pad to sizeof(struct socketaddr)
}

// generic socket address structure(for connect, bind, and accept)
struct socketaddr {
  uint16_t sa_family; // protocol family
  char sa_data[14];   // address data
}

// Steven
typedef struct socketaddr SA;
