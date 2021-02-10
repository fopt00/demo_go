package main

import "demo_go/tour"

func main() {
	//complex_types.ArrayEntry()
	//complex_types.SliceEntry()
	//complex_types.MapEntry()
	//complex_types.StructEntry()
	//complex_types.JsonEntry()
	//complex_types.TemplateEntry()

	//function.FunctionBasicEntry()
	//function.AdvancedEntry()
	//function.VariadicEntry()
	//function.DeferredEntry()
	//function.PanicEntry()

	//method.BasicEntry()
	//method.PointerReceiverEntry()
	//method.StructMethodEntry()
	//method.BitsetEntry()

	//interfaces.BasicEntry()
	//interfaces.UsingFlagEntry()
	//interfaces.FlagValueEntry()
	//interfaces.HttpServerEntry()
	//interfaces.ErrorEntry()
	//interfaces.TypeSwitchEntry()

	//goroutine.RoutineEntry()
	//goroutine.ChanEntry()
	//goroutine.PipelineEntry()
	//goroutine.OneDirectionChanEntry()
	//goroutine.ConcurrentEntry()
	//goroutine.CrawlerEntry()
	//goroutine.SpiderEntry()
	//goroutine.SelectEntry()
	//goroutine.CancelContextEntry()

	//concurrent.MutexEntry()
	//concurrent.CommunicationEntry()
	//concurrent.SyncOnceEntry()
	//concurrent.RaceConditionEntry()

	//proxy.EntryOfProxy()

	//reflect.BasicEntry()
	//reflect.KindOfEntry()
	//reflect.IteratingStructEntry()
	//reflect.JsonEntry()
	//reflect.ElemEntry()

	/*fmt.Println(tour.Add(1.0, 2.0))
	x, y := tour.Swap(1, 2)
	fmt.Println(x, y)
	a, b := tour.Split(100)
	fmt.Println(a, b)*/

	//tour.Tester4()
	//tour.ConvertTester()
	/*pi := tour.Pi
	fmt.Println(pi)*/

	/*big := tour.NeedFloat(tour.Big)
	smallFloat := tour.NeedFloat(tour.Small)
	smallInt := tour.NeedInt(tour.Small)
	fmt.Println(big, smallFloat, smallInt)*/

	/*text := "hello good good up"
	count := tour.WordCount(text)
	for k, v := range count {
		fmt.Printf("%v = %v\n", k, v)
	}*/

	/*var a = 11
	tour.DoubleValue(&a)
	fmt.Println(a)*/

	//tour.Entry()
	tour.TestLoop()
}
