//LoadBalancer with Golang//
LoadBalancer is a program responsible for keeping the track of traffic on servers, accordingly it forwards the upcoming request.Method used to create load balancer is Round Robin(route traffic between servers).
We used 'http util' to build Reverse Proxy server(redirects request and hides the location of server from client, for security).
Load Balancer is implemented through struct -> {Port,Servers,Round Robin Count}.
Two functions: (1)create a new loadbalancer (2)create a new server
Two methods: (1)get address of a server (2)get if the serverisAlive. Else go to the next server - use round robin iterator.
