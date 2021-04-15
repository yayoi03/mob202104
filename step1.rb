require 'date'

class Researve
    def initialize
        @researve_list = []
    end

    def register(start_time, finish_time)
        if finish_time==15 ||  finish_time==30 || finish_time==60
            _start = DateTime.parse(start_time)
            _end = _start + Rational(finish_time, 24*60)
            _start = _start.strftime("%T")
            _end = _end.strftime("%T")
            @researve_list.push([_start, _end])
            puts "#{_start}から#{_end}までの予約が完了しました．"
        else
            puts "予約時間は15分，30分，60分から指定してください．"
        end
    end

    def show
        for researve in @researve_list
            puts "#{researve[0]} - #{researve[1]}"
        end
    end
end

hoge = Researve.new

hoge.register('12:00:00', 15)
hoge.register('12:11:00', 12)
hoge.register('13:00:00', 30)

hoge.show