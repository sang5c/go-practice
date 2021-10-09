package main

//
//var (
//	reallyLongCalculation = func(done <-chan interface{}, value interface{}) {
//		intermediateResult := longCalculation(value)
//		select {
//		case <-done:
//			return nil
//		default:
//		}
//
//		return longCalculation(intermediateResult)
//	}
//)
//
//func main() {
//	var value interface{}
//	select {
//	case <-done:
//		return
//	case value = <-valueStream:
//	}
//
//	result := reallyLongCalculation(value)
//
//	select {
//	case <-done:
//		return
//	case resultStream <- result:
//	}
//}
