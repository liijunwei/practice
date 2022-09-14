
// ip socket address structure
struct sockaddr_in {
  uint16_t      sin_family;       // protocol family, always AF_INET
  uint16_t      sin_port;         // port number in network byte order
  struct        in_addr sin_addr; // ip address  in network byte order
  unsigned char sin_zero[8];      // pad to sizeof(struct socketaddr)
}
