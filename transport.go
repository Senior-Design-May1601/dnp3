// Transport Layer functions
/*
* Not sure what to do with this. 1 byte header 
* 1 bit fin - message end
* 1 bit fir - message sart 
* 6 bit sequence number
*/
package main

func TransportRead(data []byte)([]byte){
	return data[1:]
}
