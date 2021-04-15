#まずはある会議室の前の iPad のシステムから実装してみましょう

#好きな時刻から 15min / 30min / 60min の時間の予約を登録することができる
#想定外の時間(e.g. 10min / 90min ...)の予約が登録しようとされた場合には登録させない
#現在登録されている 予約の 開始時刻と終了時刻の組 が一覧で確認できる

# 9:00~22:00

def register(start_at, minutes)
  # start_at -> int 1200
  # minutes -> int 15, 30, 90
  counter = 0
  registered_time = {}
  if minutes.include([15, 30, 60])
    # correct
    finish_at = start_at + minutes
    registered_time[start_at] =
  else
    puts 'wrong pattern'
  end
end

