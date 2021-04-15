require 'time'

class Reservation
  def initialize
    @@reservations = {}
  end

  def add(start_period, during)
    puts "時間が適切ではありません" unless valid_time(during)
    @@reservations["#{@@reservations.length + 1}"] = {
      start_period: start_period,
      end_period: start_period + during
    }
  end

  def valid_time(time)
    return [15, 30, 60].include?(time)
  end

  def ls
    puts @@reservations
  end

  def extend_time(id:, time:)
    puts "時間が適切ではありません" unless valid_time(time)
    will_end_period = reservation["#{id}"][:end_period] + time
    results = @@reservations.select do |_, reservation|
      will_end_period > reservation[:start_period] && will_end_period < reservation[:end_period]
    end
    if results.empty?
      reservation["#{id}"][]
    end
  end
end
