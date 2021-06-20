package main

import "io"

type MyReader struct {
	r        io.Reader
	cur_num int64
	limit  int64 // limit
}


func (my *MyReader) Read(b []byte) (n int, err error) {
	n ,err =  my.r.Read(b[:my.limit])
	my.cur_num += int64(n)
	if my.cur_num >= my.limit {
		err = io.EOF
	}
	return
}


func LimitReader(r io.Reader, n int64 ) io.Reader {
	return &MyReader{r, n, 0}
}

func main()  {

}

