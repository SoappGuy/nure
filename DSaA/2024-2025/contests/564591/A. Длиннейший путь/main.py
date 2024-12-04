n, m = map(int, input().split())

edges = []
for _ in range(m):
    u, v = map(int, input().split())
    edges.append([u, v])

adj = [[] for _ in range(n + 1)]


def tsort(v, visited, order):
    visited[v] = True
    for u in adj[v]:
        if not visited[u]:
            tsort(u, visited, order)
    order.append(v)


for edge in edges:
    adj[edge[0]].append(edge[1])

visited = [False] * (n + 1)
order = []
for i in range(1, n + 1):
    if not visited[i]:
        tsort(i, visited, order)

order.reverse()

dist = [0] * (n + 1)
for u in order:
    for v in adj[u]:
        dist[v] = max(dist[v], dist[u] + 1)

print(max(dist))
