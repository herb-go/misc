# debounce 
去抖库，在一段时间内只执行一次指定函数，避免过多的无效操作

## 使用方式
    func New(duration time.Duration, callback func()) 

duration 指定的时间周期

callback :func() 真正执行的函数，需要自行处理recovery

创建好debounce后，通过exec方式调用，返回方法是否直接被调用了

    immediate:=debouce.Exec()

## 额外参数

Leading:在时间周期的开始而非结束执行原始代码

    debouce.LEading=true

MaxDuration:去抖在持续时间内被重复调用会延长去抖周期，延长的去抖周期不超过此值。默认为去抖周期的2倍

    debouce.MaxDuration=10*tome.Second

默认的倍数可以通过设置DefaultMaxDurationMagnification值来调整

    debouce.DefaultMaxDurationMagnification=int64 3