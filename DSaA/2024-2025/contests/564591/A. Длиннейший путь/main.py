n, m = map(int, input().split())
adj = [[] for _ in range(n + 1)]
for node in range(m):
    a, b = map(int, input().split())
    adj[a].append(b)


dp = [0] * (n + 1)
visited = [False] * (n + 1)


def dfs(node, adj, visited):

    visited[node] = True
    for i in range(0, len(adj[node])):
        if not visited[adj[node][i]]:
            dfs(adj[node][i], adj, visited)

        dp[node] = max(dp[node], 1 + dp[adj[node][i]])


for node in range(1, n + 1):
    if not visited[node]:
        dfs(node, adj, visited)

ans = 0

for node in range(1, n + 1):
    ans = max(ans, dp[node])

print(ans)
