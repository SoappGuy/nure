
#include <bits/stdc++.h>
#include <vector>
using namespace std;

vector<int> find_lis_reverse(vector<int> &nums) {
  int n = nums.size();
  vector<int> lis;

  for (int i = n - 1; i >= 0; i--) {
    if (lis.empty() || nums[i] < lis.back()) {
      lis.push_back(nums[i]);
    } else {
      int low = upper_bound(lis.begin(), lis.end(), nums[i]) - lis.begin();
      lis[low] = nums[i];
    }
  }
  reverse(lis.begin(), lis.end());

  return lis;
}

int main() {
  int n;
  cin >> n;

  vector<int> nums(n);
  for (int i = 0; i < n; i++) {
    cin >> nums[i];
  }

  vector<int> ans = find_lis_reverse(nums);
  cout << ans.size() << endl;
  for (int i = 0; i < ans.size(); i++) {
    cout << ans[i] << " ";
  }
  cout << endl;

  return 0;
}
