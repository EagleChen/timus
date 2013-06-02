#include <iostream>
#define __int64 long long

unsigned __int64 f(unsigned __int64 n, unsigned __int64 m)
{
  unsigned __int64 v = 1;
  for (unsigned __int64 i = 1; i <= n; i++) v = v * (m + i) / i;
  return v;
}

int main()
{
  unsigned __int64 n, a, b;
  std::cin >> n >> a >> b;
  std::cout << f(n, a) * f(n, b) << std::endl;
}
