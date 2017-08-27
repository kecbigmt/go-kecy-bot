package lib

import (
    "fmt"
    "strconv"
    "github.com/ChimeraCoder/anaconda"
    "net/url"
)

const (
    consumerKey = "cyMVZI0vVr86pyqYzDwAZJoRo"
    consumerSecret = "ue16UYLhn07o3sOYILBjMee5czwPBG9sNGbhUxNq28vEZEFLVt"
    accessToken = "1346132293-0syOFX5ENvtZbXUALB0E7hYkd4UxkRI91aRqTLn"
    accessTokenSecret = "qUnprcBMZ6dYe9c4Jjst0eXJThVHvUm8wUZlpJ7KX3M6Y"
)

type Tweets struct {
  Header []string
  Rows [][]string
}
func (t *Tweets) String() string{
  var s string
  for _, r := range t.Rows{
    s += fmt.Sprintf("[%s] (%s Likes, %s RT)\n%s\n------------------------------------------------\n", r[0], r[3], r[4], r[2])
  }
  return s
}
func NewTweets(n int) *Tweets{
  t := new(Tweets)
  t.Header = []string{
    "CreatedAt",
    "TweetId",
    "Text",
    "FavoriteCount",
    "RetweetCount",
    "Source",
    "UserId",
    "UserScreenName",
    "Name",
    "UserFriendsCount",
    "UserFollowersCount",
  }
  t.Rows = make([][]string, n)
  return t
}

func makeRow (tweet anaconda.Tweet) []string {
  row := []string{
    tweet.CreatedAt,
    tweet.IdStr,
    tweet.Text,
    strconv.Itoa(tweet.FavoriteCount),
    strconv.Itoa(tweet.RetweetCount),
    tweet.Source,
    tweet.User.IdStr,
    tweet.User.ScreenName,
    tweet.User.Name,
    strconv.Itoa(tweet.User.FriendsCount),
    strconv.Itoa(tweet.User.FollowersCount),
  }
  return row
}

func Collect(s string, c string, l string) *Tweets{
    anaconda.SetConsumerKey(consumerKey)
    anaconda.SetConsumerSecret(consumerSecret)
    api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

    loop_n := 0
    tail_c := 0
    i, _ := strconv.Atoi(c)
    if i > 100 {
      loop_n = i/100
      tail_c = i%100
      c = "100"
    }

    v := url.Values{}
    v.Set("count", c)
    if l != "" {
      v.Add("lang", l)
    }
    searchResult, _ := api.GetSearch(s, v)
    output := NewTweets(i)
    for i, tweet := range searchResult.Statuses {
      output.Rows[i] = makeRow(tweet)
    }
    for i1 := 0; i1 < loop_n; i1++ {
      searchResult, _ = searchResult.GetNext(api)
      for i2, tweet := range searchResult.Statuses {
        if i1+1 >= loop_n && i2+1 > tail_c {
          break
        }
        n := (i1+1)*100 + i2
        output.Rows[n] = makeRow(tweet)
      }
    }
    return output
}
