#include <cstdio>
#include <cstring>
#include <cmath>
#define inf 100000000.000
#define MAX 1010
double dp[2][MAX];
bool node[MAX][MAX];
int N,M,K;
const double D=sqrt(1.*100*100+1.*100*100);

void DP()
{
    for(int j=0; j<=M; j++)
        dp[0][j]=dp[1][j]=inf;
    dp[1][0]=0;

    int c=0;
    while(c<=N)
    {
        for(int j=0; j<=M; j++)
        {
            if(j-1>=0 && dp[1][j-1]+100 < dp[1][j])
                dp[1][j]=dp[1][j-1]+100;

            if(dp[0][j]+100 < dp[1][j])
                dp[1][j]=dp[0][j]+100;

            if(node[c][j] && j-1>=0 && dp[0][j-1]+D < dp[1][j])
                dp[1][j]=dp[0][j-1]+D;

        }

        for(int j=0; j<=M; j++)
        {
            dp[0][j]=dp[1][j];
            dp[1][j]=inf;
        }
        c++;
    }

    printf("%.0f\n",dp[0][M]);

    return ;
}
int main()
{
    while(scanf("%d%d",&N,&M)!=EOF)
    {
        scanf("%d",&K);
        int x,y;
        memset(node,0,sizeof(node));
        for(int i=0; i<K; i++)
        {
            scanf("%d%d",&x,&y);
            node[x][y]=1;
        }
        DP();
    }
    return 0;
}
