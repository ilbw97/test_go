package main

/*
func main() {
	sum := 1
	for sum < 1000 { //초기화 구문과 사후 구문은 필수는 아님
		sum += sum
	}
	fmt.Println(sum)
}
*/
/*func main() {
	sum := 1
	for sum < 1000 { //;을 생략할 수 있다는 점에서 C의 while == Go의 for
		sum += sum
	}
	fmt.Println(sum)
}*/

/*func main() {
	for {
	}
}*/
/*
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}*/
/*
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim { //이 라인에서 선언된 변수들은 if문의 끝까지만 존재
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	} //여기서 죽음
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
*/
/*
func Sqrt(x float64) float64 {
	z := float64(1)
	for {
		z -= z - (z*z-x)/(2*z)
	}
}

func main() {
	fmt.Println(Sqrt(2))
}
*/
/*
func main() {
	var sum = 0
	var num int
Loop1:
	for i := 0; i < 10; i++ {
	Loop2:
		for j := 0; j < 10; j++ {
			fmt.Print("정수 입력:")
			fmt.Scanln(&num)
			if num < 0 {
				break Loop1
			} else {
				if num > 100 {
					break Loop2
				}
			}
			sum += num
			fmt.Println("j:", j)
		}
		fmt.Println("i:", i)
	}
	fmt.Println("합계:", sum)
}
*/
//Switch : case
/*
func main() {
	fmt.Printf("GO runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s. \n", os)
		//freebsd, openbsdd
		//plan9, windows..
	}
}
*/

/*func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away")
	}
}
*/
/*
func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
*/
/*
func main() {
	var score int
	fmt.Print("점수 입력:")
	fmt.Scanln(&score)
	if (score < 0) || (score > 100) {
		fmt.Println("잘못 입력하셨네요.")
		return
	}
	switch score / 10 {
	case 10:
	case 9:
		fmt.Println("A")
	case 8:
		fmt.Println("B")
	case 7:
		fmt.Println("C")
	case 6:
		fmt.Println("D")
	default:
		fmt.Println("F")
	}*/
/*
	var x, y int
	fmt.Print("두 개의 정수 입력:")
	fmt.Scanln(&x, &y)

	switch {
	case x > y:
		fmt.Println("큰 값:", x)
	case x < y:
		fmt.Println("큰 값:", y)
	default:
		fmt.Println("두 수는 서로 같습니다.")
	}
}*/
//fallthrough : 특정 case문 실행 후 다음 case문 실행하고 싶을 때 사용. c의 switch에서 break 생략한것과 동일. *맨 마지막 case에는 사용 불가

/*func main() {
	i := 4
	switch i {
	case 4:
		fmt.Println("4 이상")
		fallthrough
	case 3:
		fmt.Println("3 이상")
		fallthrough
	case 2:
		fmt.Println("2 이상")
		fallthrough
	case 1:
		fmt.Println("1 이상")
		fallthrough
	case 0:
		fmt.Println("0 이상")
	}
}
*/
//case 여러개 함께 처리
/*
func main() {
	i := 3
	switch i {
	case 2, 4, 6: //i가 2,4,6일 때
		fmt.Println("짝수")
	case 1, 3, 5:
		fmt.Println("홀수")
	}
}
*/

/*
func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done!")
}
*/
/*
func main() {
	const (
		JAN = iota + 1
		FEB
		MAR
		APR
		MAY
		JUN
		JUL
		AUG
		SEP
		OCT
		NOV
		DEC
	) //월
	fmt.Println("JAN:", JAN)
	fmt.Println("JUL:", JUL)
	fmt.Println("DEC:", DEC)

	const (
		c0 = iota * 10
		c1 = iota * 10
		c2 = iota * 10
		c3 = iota * 10
		c4 = iota * 10
	) //10씩 증가하는 상수
	fmt.Println("c1 : ", c1)
	fmt.Println("c1 : ", c2)
	fmt.Println("c1 : ", c3)
	fmt.Println("c1 : ", c4)

}
*/
/*
func main() {
	const max_elem int = 5
	var i int = 0
	fmt.Println("i:", i)
	i = (i + 1) % max_elem
	fmt.Println("i:", i)
	i = (i + 1) % max_elem
	fmt.Println("i:", i)
	i = (i + 1) % max_elem
	fmt.Println("i:", i)
	i = (i + 1) % max_elem
	fmt.Println("i:", i)
	i = (i + 1) % max_elem
	fmt.Println("i:", i)
	i = (i + 1) % max_elem
	fmt.Println("i:", i)
}
*/
/*
func main() {
	fmt.Printf("2<3=%b\n", 2 < 3)          //bool
	fmt.Printf("23은 이진수로 %b\n", 23)        //이진수
	fmt.Printf("family name %c\n", '홍')    //문자
	fmt.Printf("10진수 출력:%d\n", 23)         //10진수
	fmt.Printf("8진수 출력:%o\n", 23)          //8진수
	fmt.Printf("16진수 출력:%x\n", 23)         //16진수,a~f
	fmt.Printf("16진수 출력:%X\n", 23)         //16진수,A~F
	fmt.Printf("고정소수점:%f\n", 123.4567)     //고정소수점
	fmt.Printf("고정소수점:%F\n", 123.4567)     //고정소수점
	fmt.Printf("지수 표현:%e\n", 123.4567)     //지수 표현, e
	fmt.Printf("지수 표현:%E\n", 123.4567)     //지수 표현, E
	fmt.Printf("간단한 실수 표현:%g\n", 123.4567) //간단한 실수 표현
	fmt.Printf("간단한 실수 표현:%g\n", 123.4567) //간단한 실수 표현
	fmt.Printf("문자열:%s", "홍길동\n")          //문자열
	var a int = 32
	fmt.Printf("메모리주소:%p\n", a)                     //포인터
	fmt.Printf("유니코드 %U\n", '\ud55c')               //유니코드
	fmt.Printf("%T\n", 23)                          //타입
	fmt.Printf("모든 형식:%v , %v\n", 23, 'a')          //모든 형식
	fmt.Printf("형식도 함께:%d %#o , %#x\n", 23, 23, 23) //%#v 형식을 구분할 수 있게
	fmt.Printf("%4d %04d\n", 23, 23)                //출력 폭 지정, 빈 곳 0출력
	fmt.Printf("%-4d %-4d\n", 23, 23)               //왼쪽 정렬
	fmt.Printf("%9.2f\n", 123.4567)                 //소수점 이하 자리 출력 지
}
*/
/*
func main() {
	var name string
	fmt.Print("이름:")
	fmt.Scanln(&name)
	var num int
	fmt.Print("번호:")
	fmt.Scanln(&num)
	var addr string
	fmt.Print("주소:")
	fmt.Scanln(&addr)
	fmt.Println("이름은 ", name, "번호는 ", num)
	fmt.Println("주소는 ", addr)

	var name, addr string
	var num int
	fmt.Println("이름 번호 주소")
	var re int
	re, _ = fmt.Scanln(&name, &num, &addr)
	fmt.Println("이름은 ", name, " 번호는 ", num)
	fmt.Println("주소는 ", addr)
	fmt.Println(re)
*/
/*
		var n1, n2, n3, n4 int
		fmt.Println("IPv4 주소")
		fmt.Scanf("%d.%d.%d.%d", &n1, &n2, &n3, &n4)
		fmt.Printf("입력한 IPv4주소는 %d.%d.%d.%d\n", n1, n2, n3, n4)

}*/
/*
func main() {
	var scores [3]int = [3]int{1, 2}
	fmt.Println("=== scores ===")
	for i := 0; i < len(scores); i++ {
		fmt.Println(scores[i])
	}
	var scores2 = [3]int{1, 2}
	fmt.Println("=== scores2 ===")
	for i := 0; i < len(scores2); i++ {
		fmt.Println(scores2[i])
	}
	var scores3 = [3]int{1, 2}
	fmt.Println("=== scores3 ===")
	for i := 0; i < len(scores3); i++ {
		fmt.Println(scores3[i])
	}
	scores4 := [...]int{1, 2}
	fmt.Println("=== scores4 ===")
	for i := 0; i < len(scores4); i++ {
		fmt.Println(scores4[i])
	}
}
*/
/*
func main() {
	arr := [5]int{12, 34, 23, 56, 34}
	//var length int
	//length = len(arr)
	//for i := 0; i < length; i++ {
	//	fmt.Println(i, ":", arr[i])
	//}
	for _, i := range arr {
		//fmt.Println(i, ":", value)
		fmt.Println(i)
	}
}
*/
/*
func main() {
	var n int
	fmt.Print("학생수:")
	fmt.Scanln(&n)

	var scores []int = make([]int, n)
	var i int
	for i = 0; i < n; i++ {
		fmt.Printf("%d번 성적 : ", i+1)
		fmt.Scanln(&scores[i])
	}
	fmt.Println("==== 학생 성적 출력 ====")

	var sum int
	for i = 0; i < n; i++ {
		sum += scores[i]
		fmt.Printf("%d번 성적: %d\n", i+1, sum)
	}
	fmt.Printf("총점:%d\n", sum)
}
*/
/*
func main() {
	var arr [5]int = [5]int{90, 88, 76, 80, 99}
	var scores []int = make([]int, 5, 10)
	var i int

	for i = 0; i < 5; i++ {
		scores[i] = arr[i]
	}
	fmt.Printf("저장소 크기:%d 보관한 자료 개수:%d\n", cap(scores), len(scores))
	var score int
	for i = 5; i < 10; i++ {
		fmt.Printf("%d 번 성적 : ", i+1)
		fmt.Scanln(&score)
		scores = append(scores, score)
		fmt.Printf("저장소 크기:%d 보관한 자료 개수:%d\n", cap(scores), len(scores))
	}
	fmt.Println("==== 학생 성적 출력 ====")

	var sum int
	for i = 0; i < 10; i++ {
		sum += scores[i]
		fmt.Printf("%d번 성적:%d\n", i+1, sum)
	}
	fmt.Printf("총점:%d\n", sum)
}
*/
/*
func main() {
	var datas []int
	fmt.Println(datas)
	datas = append(datas, 9, 8, 7, 6, 5, 4)
	fmt.Println(datas)
}
*/
/*
func main() {
	var s []int
	fmt.Printf("용량:%d 원소 개수:%d\n", cap(s), len(s))
	var i int
	for i = 0; i < 10; i++ {
		s = append(s, i+1)
		fmt.Printf("용량:%d 원소 개수:%d\n", cap(s), len(s))
	}
}
*/
/*
func main() {
	var scores []int = []int{90, 88, 76, 80, 99}
	var i int

	fmt.Printf("저장소 크기:%d 보관한 자료 개수:%d\n", cap(scores), len(scores))
	var score int
	for i = 5; i < 10; i++ {
		fmt.Printf("%d번 성적:", i+1)
		fmt.Scanln(&score)
		scores = append(scores, score)
		fmt.Printf("저장소 크기:%d 보관한 자료 개수:%d\n", cap(scores), len(scores))
	}
	fmt.Println("==== 학생 성적 출력 ====")

	var sum int
	for i = 0; i < 10; i++ {
		sum += scores[i]
		fmt.Printf("%d번 성적:%d\n", i+1, sum)
	}
	fmt.Printf("총점:%d\n", sum)
}
*/
/*
func main() {
	//array
	var arr1 [3]int = [3]int{1, 2, 3}
	var arr2 [3]int

	fmt.Print("before arr1 : ")
	fmt.Println(arr1)
	fmt.Print("before arr2 : ")
	fmt.Println(arr2)

	arr2 = arr1

	arr2[0] = 8
	fmt.Print("after arr1 : ")
	fmt.Println(arr1)
	fmt.Print("after arr2 : ")
	fmt.Println(arr2)

	//slice
	var s1 []int = []int{1, 2, 3}
	var s2 []int

	fmt.Print("before s1 : ")
	fmt.Println(s1)
	fmt.Print("before s2 : ")
	fmt.Println(s2)

	s2 = s1

	s2[0] = 8
	fmt.Print("after s1 : ")
	fmt.Println(s1)
	fmt.Print("after s2 : ")
	fmt.Println(s2)
}
*/
/*
func main() {
	var s1 []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var s2 []int = make([]int, 3, 5)

	s2[0] = 8
	s2[1] = 9
	fmt.Printf("s1 용량 %d 원소 개수:%d\n", cap(s1), len(s1))
	fmt.Println(s1)
	fmt.Printf("s2 용량 %d 원소 개수:%d\n", cap(s2), len(s2))
	fmt.Println(s2)

	copy(s2, s1)

	fmt.Println("===copy(s2,s1) 수행 후===")
	fmt.Printf("s1 용량 %d 원소 개수:%d\n", cap(s1), len(s1))
	fmt.Println(s1)
	fmt.Printf("s2 용량 %d 원소 개수:%d\n", cap(s2), len(s2))
	fmt.Println(s2)

}
*/
/*
func main() {
	var origin_s []int = []int{10, 23, 34, 47, 62, 7, 89, 91, 102}
	var start int
	var last int

	fmt.Print("원본 슬라이스:")
	fmt.Println(origin_s)

	fmt.Println("원본 슬라이스에서 부분 슬라이스 만들기")
	fmt.Printf("시작 인덱스:")
	fmt.Scanln(&start)
	fmt.Printf("끝 인덱스:")
	fmt.Scanln(&last)

	var sub_s []int = origin_s[start : last+1]
	fmt.Print("부분 슬라이스:")
	fmt.Println(sub_s)
}
*/
/*
func main() {
	i, j := 42, 2701

	p := &i         //i를 가리키는 pointer
	fmt.Println(*p) //pointer를 통해 i의 값 읽기
	*p = 21         //pointer를 통해 i값 설정
	fmt.Println(i)

	p = &j       //j를 가리키는 pointer
	*p = *p / 37 //pointer를 통해 j를 나눔
	fmt.Println(j)
}
*/
/*
type Vertex struct {
	X int
	Y int
}

//구조체 인스턴스 언언 방법
var (
	//1. 일반적인 방식. X=1, Y=2로 초기화
	v1 = Vertex{1, 2}
	//2. X값만 지정, Y는 int에 zero value로 설정
	v2 = Vertex{X: 1}
	//3. X,Y 모두 int에 zero value로 설정
	v3 = Vertex{}
)

func main() {
	fmt.Println("v1.X값:", v1.X)
	v1.X = 4
	fmt.Println("v1.X = 4로 바꾼 v1.X의 값 : ", v1.X)

	//4. 구조체 pointer로도 구조체의 값을 바꿀 수 있다.
	var p = &v1
	p.X = 10
	fmt.Println("pointer로 바꾼 v1.X값 : ", v1.X)
}
*/
/*
func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println("배열 names:", names)

	fmt.Println("1. 슬라이스 선언")
	// 슬라이스 선언방법
	// 1. 일반적인 선언방법 : 변수 선언과 비슷합니다. 슬라이스타입은 []type입니다.
	var s1 []string = names[0:3]
	// 2. 슬라이스도 var키워드와 타입 명시를 생략할 수 있습니다.
	s2 := names[0:2]

	fmt.Println("names[0:3]:", s1)
	fmt.Println("names[0:2]:", s2)

	//s1에서 값을 바꾸면 names, s1에서도 바뀐 값을 볼 수 있습니다.
	fmt.Println("2. 슬라이스로 값 변경")
	fmt.Println("s1[0]", s1[0])
	s1[0] = "XXX"
	fmt.Println("s1[0] = XXX 실행 후 s1:", s1)
	fmt.Println("s1[0] = XXX 실행 후 s2:", s2)
	fmt.Println("s1[0] = XXX 실행 후 names:", names)

	s2 = s1[0:2]
	fmt.Println("s2 = s1[0:2] 실행 후  s2:", s2)
}
*/
/*
func main() {
	fmt.Println("1. 슬라이스 리터럴 선언")
	//1. 기본형 슬라이스 리터럴
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("기본형 슬라이스 리터럴 : ", q)
	//2. 구조체 슬라이스 리터럴
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println("구조체 슬라이스 리터럴:", s)

	fmt.Println("2. 슬라이스를 슬라이스")
	q = q[:2]
	fmt.Println("q[:2]:", q)

	q = q[1:]
	fmt.Println("q[1:]:", q)
}
*/
/*
func main() {
	//make로 가변 길이 만들기
	a := make([]int, 5)
	fmt.Printf("a := make([]int, 5)의 \t")
	printSlice(a)

	b := make([]int, 0, 5)
	fmt.Printf("b := make([]int, 0, 5)의 \t")
	printSlice(b)

	c := b[:2]
	fmt.Printf("c := b[:2]의 \t")
	printSlice(c)

	d := c[2:5]
	fmt.Printf("d := c[2:5]의 \t")
	printSlice(d)

	//한번에 여러개 원소 추가 가능
	d = append(d, 1, 2, 3)
	fmt.Printf("d = append(d, 1,2,3)후\t")
	printSlice(d)
}
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
*/
/*
type Vertex struct {
	Lat, Long float64
}

func main() {
	//1. map 사용
	//map[string] type 변수 선언
	var mymap map[string]Vertex
	//make()로 맵 생성
	mymap = make(map[string]Vertex)
	mymap["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println("1. mymap[\"Bell Labs\"]: ", mymap["Bell Labs"])
	//2. map literal 사용
	var mymap_literal = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println("2. mymap_literal[\"Google\"]", mymap_literal["Google"])
}
*/
/*
func main() {
	m := make(map[string]int)

	//1. key-value 지정
	m["Answer"] = 42
	fmt.Println("m[\"Answer\"]값은:", m["Answer"])

	//2. key-value 삭제
	delete(m, "Answer")
	fmt.Println("m[\"Answer\"]값은:", m["Answer"])

	//3. key 존재 확인
	v, ok := m["Answer"]
	fmt.Println("m[\"Answer\"]값은", v, "존재하나요?", ok)
}
*/
/*
func main() {
	solarSystem := make(map[string]float32) //key : string, value : float32인 map 생성, 공간 할다으

	solarSystem["Mercury"] = 87.969 // 맵[키] = 값
	solarSystem["Venus"] = 224.70069
	solarSystem["Earth"] = 365.25641
	solarSystem["Mars"] = 686.9600
	solarSystem["Jupiter"] = 4333.2867
	solarSystem["Saturn"] = 10756.1995
	solarSystem["Uranus"] = 30707.4896
	solarSystem["Neptune"] = 60223.3528

	value, ok := solarSystem["Pluto"] // map에 key가 있는지 검사할 때는 return 값을 두 개 사용.
	if ok == false {
		fmt.Println("Not exist")
	} else {
		fmt.Println(value, ok)

	}

	if value, ok := solarSystem["Saturn"]; ok {
		fmt.Println(value)
	}
	for key, value := range solarSystem {
		fmt.Println(key, ":", value)
	}
	for _, value := range solarSystem {
		fmt.Println(value)
	}
}
*/
//func 함수명(매개변수명 자료형) (리턴값_변수명 자료형){}
/*
func sum(a int, b int) (r int) { //return 값 변수 이름을 r로 지정
	r = a + b
	return //return 값 변수 사용시 return 뒤에 변수 지정 x
}
func main() {
	r := sum(1, 2)
	fmt.Println(r)
}
*/
//가변인자 사용하기
//func 함수명(매개변수명 ...자료형) 리턴값 _자료형
/*
func sum(n ...int) int { //int형 가변인자를 받는 함수 정의
	//여기서 가변인자 함수는 INT형 값만 여러 개 받도록 되어있고, SLICE 자체는 받을 수 없다.
	//따라서 매개변수에 슬라이스만 넣지 않고 뒤에 ...을 붙인다.
	//...을 붙이면 슬라이스에 들어있는 요소를 각각 넘겨준다.
	total := 0
	for _, value := range n { //range로 가변인자의 모든 값을 꺼냄
		total += value //꺼낸 값 모두 더하기
	}
	return total
}
func main() {
	//r := sum(1, 2, 3, 4, 5)
	//fmt.Println(r)
	n := []int{1, 2, 3, 4, 5} //slice 선언
	r := sum(n...)            // ...를 사용해 가변인자에 슬라이스를 바로 넘겨줌

	fmt.Println(r)
}
*/
//재귀호출 사용하기
/*
func factorial(n uint64) uint64 {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
func main() {
	fmt.Println(factorial(5))
}
*/
//함수를 변수에 저장하기
//var 변수명 func(매개변수명 자료형) 리턴값_자료형 = 함수명
/*
func sum(a int, b int) int {
	return a + b
}
func main() {
	var hello func(a int, b int) int = sum //함수를 저장하는 변수 선언하고 함수 대입
	world := sum                           //변수 선언과 동시에 함수를 바로 대입

	fmt.Println(hello(1, 2))
	fmt.Println(world(1, 2))
}
*/
//슬라이스 = []func(매개변수명 자료형) 리턴값_자료형{함수명1, 함수명2}
/*
func sum(a int, b int) int {
	return a + b
}
func diff(a int, b int) int {
	return a - b
}
func main() {
	f := []func(int, int) int{sum, diff} //함수를 저장할 수 있는 slice 생성한 뒤 함수로 초기화

	fmt.Println(f[0](1, 2))
	fmt.Println(f[1](1, 2))
}
*/
//맵 := map[키_자료형]func(매개변수명 자료형) 리턴값_자료형{"키":함수명}
/*
func sum(a int, b int) int {
	return a + b
}
func diff(a int, b int) int {
	return a - b
}
func main() {
	f := map[string]func(int, int) int{ //함수를 저장할 수 있는 맵을 생성한 뒤 함수로 초기화
		"sum":  sum,
		"diff": diff,
	}
	fmt.Println(f["sum"](1, 2))  //3 : map에 sum 키를 지정해 함수 호출
	fmt.Println(f["diff"](1, 2)) //-1 : map에 diff 키를 지정해 함수 호출
}
*/
//익명 함수 사용하기
//func(매개변수명 자료형) 리턴값_자료형 {}()
/*
func main() {
	func() { //함수에 이름이 없음
		fmt.Println("Hello world!")
	}()
	func(s string) { //익명 함수를 정의한 뒤
		fmt.Println(s)
	}("Hello World!") //바로 호출
	r := func(a int, b int) int { //익명 함수를 정의한 뒤
		return a + b
	}(1, 2) //바로 호출해 return 값을 변수 r에 저장
	fmt.Println(r) //3
}
*/
//Closure : 함수 안에서 함수를 선언 및 정의할 수 있고, 바깥쪽 함수에 선언된 변수에도 접근할 수 있는 함수.
//변수 := func(매개변수명 자료형) 리턴값_자료형 {}
/*func main() { //함수 안에서
	sum := func(a, b int) int { //익명 함수를 선언 및 정의
		return a + b
	}
	r := sum(1, 2) //익명함수 사용
	fmt.Println(r)//3

	a, b := 3, 5
	f := func(x int) int {
		return a*x + b //함수 바깥의 변수 a,b 사용
	}
	y := f(5)
	fmt.Println(y) //20
}
*/
//클로저를 사용하는 이유
/*
func calc() func(x int) int {
	a, b := 3, 5
	return func(x int) int {
		return a*x + b //클로저이므로 함수를 호출 할 때 마다 변수 a와 b의 값을 사용할 수 있음
	}
	//익명 함수를 리턴
}
func main() {
	f := calc() // calc 함수를 실행해 리턴값으로 나온 클로저를 변수에 저장. 원래 calc()의 지역변수 a,b는 f에 저장과 동시에 소멸된다.
	i := 1
	for i <= 5 {
		fmt.Println(f(i)) //클로저를 사용하면 지역변수가 소멸되지 않고, 함수를 호출할 때마다 계속 사용할 수 있다. 클로저는 함수가 선언될 때의 환경을 계속 유지한다.
		i++
	}
}
*/
//defer 함수명() --> 지연호출
//defer 함수명(매개변수)
/*
func helloworld() {
	defer func() { //helloworld()가 끝나기 직전에 호출
		fmt.Println("world")
	}()
	func() {
		fmt.Println("Hello")
	}()
}
func ReadHello() {
	file, err := os.Open("hello.txt")
	defer file.Close() //지연 호ㅗ출한 file.Close가 맨 마지막에 호출됨

	if err != nil {
		fmt.Println(err)
		return //file.Close()호출
	}
	buf := make([]byte, 100)
	if _, err = file.Read(buf); err != nil {
		fmt.Println(err)
		return //file.Close() 호출
	}
	fmt.Println(string(buf))
	//file.Close()호츌
}
func main() {
	helloworld()
	//defer는 stack과 동일하다(LIFO) --> 맨 나중에 지연 호출한 함수가 먼저 실행됨.
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
	ReadHello()

}
*/
//panic and recover
/*func main() {
	a := [...]int{1, 2}

	for i := 0; i < 3; i++ {
		fmt.Println(a[i]) //index out of range
	}
}
*/
//panic(에러_메시지) 사용시 사용자가 에러 발생 시킬 수 있음
/*func main() {
	panic("Error !!")
	fmt.Println("Hello, world!") //실행되지 않음.
}
*/
//변수 := recover()
//recover()은 panic에서 설정한 error message를 받아올 수 있음.
/*
func f() {
	defer func() { //recover()은 defer로 사용해야 함. 그렇지 않으면 프로그램이 바로 종료됨.
		s := recover() //패닉이 발생해도 프로그램을 종료하지 않음.
		fmt.Println(s)
	}()
	a := [...]int{1, 2}
	for i := 0; i < 5; i++ {
		fmt.Println(a[i])
	}
}
func main() {
	f()

	fmt.Println("Hello World!") // 패닉이 발생했지만 계속 실행됨
}
*/
/*
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	//1. 함수를 변수에 할당해, 변수를 함수처럼 씁니다.
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println("1. 변수를 통해 함수 호출", hypot(5, 12))

	//2. 함수를 compute 함수에 인자로 전달한다.
	fmt.Println("2. 함수를 함수에 인자로 전달")
	fmt.Println("comput(hypot):\t\t", compute(hypot))
	fmt.Println("comput(math.Pow):\t", compute(math.Pow))
}
*/
/*
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
func main() {
	//adder는 클로져를 리턴, 클로져 pos, neg는 서로 다른 변수 sum을 갖음
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			i, ":",
			pos(i),
			neg(-2*i),
		)
	}
}
*/
/*
type Vertex struct {
	X, Y float64
}

//1. abs method는 리시버 인자로 v Vertex를 받는다.
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//2. 기본형 타입(여기는 float64)도 메소드를 만들 수 있습니다.
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

//3. MyFloat이 포인터가 아닌 리시버 인자이다.
func (f MyFloat) power10() {
	f = f * MyFloat(10)
}

//4. MyFloat이 포인터 리시버 인자이다.
func (f *MyFloat) power100() {
	*f = *f * MyFloat(100)
}
func main() {
	v := Vertex{3, 4}
	fmt.Println("1. 점을 씩어 메소드에 접근한다.")
	fmt.Println("v.Abs():", v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println("2. numeric type도 메소드 정의가 가능하다.")
	fmt.Println("f.Abs():", f.Abs())

	fmt.Println("3. pointer 리시버를 쓰면 메소드 내부에서 값을 바꿀 수 있다.")
	fmt.Println("기존의 f\t\t\t\t", f)
	f.power10()
	fmt.Println("일반 리시버를 써서 10을 곱한 경우\t", f)

	f.power100()
	fmt.Println("포인터 리시버를 써서 100을 곱한 경우\t", f)
}
*/
/*
type Vertex struct {
	X, Y float64
}

//1. Vertex pointer 리시버가 있는 메소드이다. Vertex or Vertex pointer로 접근 가능하다.
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//2. Vertex pointer 인자가 있는 함수다.
//Vertex pointer만 인자로 들어올 수 있다.
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)
	fmt.Println("1. Vertex{3,4}로 접근했을 떄:", v)

	p := &Vertex{3, 4}
	p.Scale(2)
	ScaleFunc(p, 10)
	fmt.Println("2. &Vertex{3,4}로 접근했을 때:", p)
}
*/
/*
type I interface {
	M()
}
type T struct {
	S string
}

//1. 별도의 키워드를 쓰지 않아도 T가 인터페이스 I를 구현하게 된다.
func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

//2. 별도의 키워드를 쓰지 않아도 F가 인터페이스 I를 구현하게 된다.
func (f F) M() {
	fmt.Println(f)
}
func main() {
	var i I

	fmt.Println("1. i = &T{\"Hello\"}에 대해")
	i = &T{"Hello"}
	describe(i)
	i.M()

	fmt.Println("2. i = F(math.Pi)")
	i = F(math.Pi)
	describe(i)
	i.M()
}
func describe(i I) {
	fmt.Printf("인터페이스의 (값, 타입) : (%v, %T)\n", i, i)
}
*/
/*
func main() {
	fmt.Println("1. empty interface에 대해")
	var i interface{}
	describe(i)

	fmt.Println("1. i = 42에 대해")
	i = 42
	describe(i)
	fmt.Println("1. i = \"hello\"에 대해")
	i = "hello"
	describe(i)
}
func describe(i interface{}) {
	fmt.Printf("인터페이스 i의 (값, 타입) : (%v, %T)\n", i, i)
}
*/
/*
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
func main() {
	go say("2. 다른 루틴")
	say("1. 이 루틴")
}
*/

//--------------------------------------------------------------
//json to map
/*
func main() {
	doc := `
	{
		"u_id" : "byeoungwoolee",
		"u_pw": "1234"
	}
	`

	var data map[string]interface{} //JSON 문서의 데이터를 저장할 공간을 map으로 선언.

	json.Unmarshal([]byte(doc), &data) //doc을 바이트 슬라이스로 변환하여 넣고,
	//data의 포인터를 넣어줌.
	fmt.Println(data["u_id"], data["u_pw"]) //byeoungwoolee 1234 : map에 key를 지정해 값을 가져옴
}
*/
//map to json
/*
func main() {
	data := make(map[string]interface{}) //문자열을 key로 하고, 모든 자료형을 저장할 수 있는 map 생성

	data["u_id"] = "byeoungwoolee"
	data["u_pw"] = "1234"

	//doc, _ := json.Marshal(data) //map을 JSON 문서로 변환 --> json.Marshal()은 []byte형식으로 리턴됨.
	doc, _ := json.MarshalIndent(data, "", "	") //JSON 문서로 변환 시 한 줄로 붙어서 나오면 읽기 힘들기 때문에
	//json.MarshalIndent()를 사용.
	//json.MarshalIndent(JSON 문서로 만들 데이터, JSON 문서의 첫 칸에 표시할 문자열.(보통 빈 문자 지정), 들여쓰기 할 문자(공백문자나 탭 문자))
	fmt.Println(string(doc)) // {"u_id":"byeoungwoolee", "u_pw":"1234"} : []byte형식을 string으로 변환해서 출력

}
*/
//json struct
//json to data
/*
type Author struct {
	Name  string
	Email string
}
type Comment struct {
	ID      uint64
	Author  Author //Author 구조체
	Content string
}
type Article struct { //JSON 은 필드 안에 다시 Object나 배열이 들어가는 방식이므로 이에 맞추어 다른 구조체를 넣어준다.
	//하위 객체면 구조체를 그대로 넣고, 배열이라면 구조체를 배열로 만든다.
	ID         uint64
	Title      string
	Author     Author //Author 구조체
	Content    string
	Recommends []string  //문자열 배열
	Comments   []Comment //Comment 구조체 배열
}

func main() {
	doc := `
	[{
		"ID": 1,
		"Title": "Hello, World",
		"Author":{
			"Name": "Maria",
			"Email": "maria@example.com"
		},
		"Content": "Hello~",
		"Recommends": [
			"John",
			"Andrew"
		],
		"Comments": [{
			"id":1,
			"Author":{
				"Name": "Andrew",
				"Email": "andrew@hello.com"
			},
			"Content": "Hello Maria"
		}]
	}]
	`
	var data []Article                 //JSON 문서의 DATA를 저장할 구조체 slice 선언
	json.Unmarshal([]byte(doc), &data) //doc의 내용을 변환하여 data에 저장
	fmt.Println(data)                  //[{1 Hello, world! {Maria maria@example.com....}}]
}
*/
//data to json
//*구조체 필드가 대문자로 시작하면 --> json 문서 안의 키도 대문자로 시작
//JSON 문서 안의 키를 소문자로 시작하고 싶다면 --> 구조체 필드에 태그 지정.
/*
type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Comment struct {
	Id      uint64 `json:"id"`
	Author  Author `json:"author"` //Author 구조체
	Content string `json:"content"`
}

type Article struct {
	Id         uint64    `json:"id"`
	Title      string    `json:"title"`
	Author     Author    `json:"author"` //Author 구조체
	Content    string    `json:"content"`
	Recommends []string  `json:"recommends"` //문자열 배열
	Comments   []Comment `json:"comments"`   //comment 구조체 배열
}

func main() {
	data := make([]Article, 1)

	data[0].Id = 1
	data[0].Title = "Hello, World!"
	data[0].Author.Name = "Maria"
	data[0].Author.Email = "maria@example.com"
	data[0].Content = "Hello~"
	data[0].Recommends = []string{"John", "Andrew"}
	data[0].Comments = make([]Comment, 1) //게시물 데이터를 저장할 슬라이스 선언 후, make함수로 공간 할당.
	//data 채워넣기
	data[0].Comments[0].Id = 1
	data[0].Comments[0].Author.Name = "Andrew"
	data[0].Comments[0].Author.Email = "andrew@hello.com"
	data[0].Comments[0].Content = "Hello Maria"

	doc, _ := json.MarshalIndent(data, "", "  ") //data를 JSON 문서로 변환

	fmt.Println(string(doc)) //[{"Id":1,"Title":"Hello, World!",...}]
}
*/
//json file 저장하기
/*
type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Comment struct {
	Id      uint64 `json:"id"`
	Author  Author `json:"author"` //Author 구조체
	Content string `json:"content"`
}

type Article struct {
	Id         uint64    `json:"id"`
	Title      string    `json:"title"`
	Author     Author    `json:"author"` //Author 구조체
	Content    string    `json:"content"`
	Recommends []string  `json:"recommends"` //문자열 배열
	Comments   []Comment `json:"comments"`   //comment 구조체 배열
}

func main() {
	data := make([]Article, 1) //값을 저장할 구조체 슬라이스 생성

	data[0].Id = 1
	data[0].Title = "Hello, World!"
	data[0].Author.Name = "Maria"
	data[0].Author.Email = "maria@example.com"
	data[0].Content = "Hello~"
	data[0].Recommends = []string{"John", "Andrew"}
	data[0].Comments = make([]Comment, 1) //게시물 데이터를 저장할 슬라이스 선언 후, make함수로 공간 할당.
	//data 채워넣기
	data[0].Comments[0].Id = 1
	data[0].Comments[0].Author.Name = "Andrew"
	data[0].Comments[0].Author.Email = "andrew@hello.com"
	data[0].Comments[0].Content = "Hello Maria"

	doc, _ := json.MarshalIndent(data, "", "  ") //data를 JSON 문서로 변환

	//ioutile.WriteFile(파일 명, json문서, unix/linux type의 file permission(os.FileMode))
	err := ioutil.WriteFile("./articles.json", doc, os.FileMode(0644)) //articles.json file에 JSON 문서 저장
	if err != nil {
		fmt.Println(err)
		return
	}
}
*/
//저장된 json file 읽어오기
/*
type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Comment struct {
	Id      uint64 `json:"id"`
	Author  Author `json:"author"` //Author 구조체
	Content string `json:"content"`
}

type Article struct {
	Id         uint64    `json:"id"`
	Title      string    `json:"title"`
	Author     Author    `json:"author"` //Author 구조체
	Content    string    `json:"content"`
	Recommends []string  `json:"recommends"` //문자열 배열
	Comments   []Comment `json:"comments"`   //comment 구조체 배열
}

func main() {
	b, err := ioutil.ReadFile("./articles.json") //articles.json file의 내용르 읽어서 byte slice에 저장
	if err != nil {
		fmt.Println(err)
		return
	}

	var data []Article       //JSON 문서의 데이터를 저장할 구조체 슬라이스 선언
	json.Unmarshal(b, &data) //JSON 문서의 내용을 변환해 data에 저장.
	fmt.Println(data)        //[{1 Hello, world! {Maria maria@exa..}}]
}
*/
//암호화 사용하기
//Hash algorithm --> MD5, SHA1, SHA256, SHA512등 DATA에서 고유한 HASH값 추출. pw저장시 사용
//func New() hash.Hash : SHA 512 해시 인스턴스 생성
//func Sum512(data []byte) [Size]byte : SHA512 해시를 계산해 리턴
//func (d *digest) Write(p []byte) (nn int, err error) : 해시 인스턴스에 데이터 추가
//func (d0 *digest) Sum(in []byte) []byte : 해시 인스턴스에 저장된 데이터의 SHA512 해시 값 추출

//SHA512 Algorithm을 사용해 data에서 hash value 추출
/*
func main() {
	s := "Hello World!"
	h1 := sha512.Sum512([]byte(s))
	fmt.Printf("%x\n", h1)

	sha := sha512.New()          //SHA512 HASH INSTANCE 생성
	sha.Write([]byte("Hello, ")) //HASH INSTANCE에 DATA 추가
	sha.Write([]byte("world!"))  //HASH INSTANCE에 DATA 추가
	h2 := sha.Sum(nil)           //HASH INSTANCE에 저장된 DATA의 SHA512 HASH 값 추출
	fmt.Printf("%x\n", h2)
}
*/
//AES 대칭키 알고리즘 사용하기 crypto/aes
//AES는 BLOCK 암호화 알고리즘이므로 KEY와 DATA의 크기가 일정해햐 한다.
//길이가 긴 데이터는 16byte씩 잘라서 암호화함. --> ECB(Electronic Codebook)
/*
func main() {
	key := "Hello, key 12345" //16byte
	s := "Hello, world! 12"   //16byte

	block, err := aes.NewCipher([]byte(key)) //AES 대칭키 암호화 블럭 생성
	if err != nil {
		fmt.Println(err)
		return //암호화 블록(cipher.Block)이 return.
	}
	ciphertext := make([]byte, len(s))
	block.Encrypt(ciphertext, []byte(s)) //평문을 AES 알고리즘으로 암호화. Encrypt(암호화된_데이터를_저장할_슬라이스, data)
	fmt.Printf("%x\n", ciphertext)

	plaintext := make([]byte, len(s))
	block.Decrypt(plaintext, ciphertext) //AES 알고리즘으로 암호화된 데이터를 평문으로 복호화
	fmt.Println(string(plaintext))
}
*/
//CBC(Cipher Block Chaining) : 긴 데이터를 안전하게 암호화하기 위해 사용.

//---------------------------------------------------------------------
/*
func Listen(net, laddr string) (Listener, error): 프로토콜, IP 주소, 포트 번호를 설정하여 네트워크 연결을 대기합니다.
func (l *TCPListener) Accept() (Conn, error): 클라이언트가 연결되면 TCP 연결(커넥션)을 리턴
func (l *TCPListener) Close() error: TCP 연결 대기를 닫음
func (c *TCPConn) Read(b []byte) (int, error): 받은 데이터를 읽음
func (c *TCPConn) Write(b []byte) (int, error): 데이터를 보냄
func (c *TCPConn) Close() error: TCP 연결을 닫음
*/
/*
func requestHandler(c net.Conn) { //client에서 받은 패킷을 처리하고, client로 패킷을 보내는 함수.
	data := make([]byte, 4096) //4096 크기의 byte slice 생성

	for { //무한 loop 을 돌면서 client에서 보낸 data를 읽어서 다시 client로 보내는 구조.
		n, err := c.Read(data) //TCP 연결 'c'에서 Read()함수로 client에서 보낸 data 읽음
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(data[:n])) //data 출력

		_, err = c.Write(data[:n]) //TCP 연결 'c'에서 Write()함수를 사용해 client로 data를 본냄.
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
func main() {
	ln, err := net.Listen("tcp", ":8000") //net.Listen()함수에 tcp와 포트번호설정해 TCP 프로토콜에 8000 포트로 연결 받음.
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close() //tcp 연결 대기 'ln'은 서버가 끝나면 닫음.

	for { // 무한 loop을 돌면서 client의 연결을 기다림.
		conn, err := ln.Accept() //client가 연결되면 TCP 연결(conn) 리턴
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer conn.Close() //main함수가 끝나기 직전에 TCP 연결을 닫음.

		go requestHandler(conn) // 그 뒤 패킷을 처리할 requestHandler()함수를 고루틴으로 ㅗ실행.
	}
}

*/
/*
func hello(n *int) {
	*n = 2
}
func main() {
	var n int = 1
	hello(&n)
	fmt.Println(n)
}
*/
/*type Rectangle struct {
	width, height int
}

func main() {
	var rect Rectangle
	var rect1 *Rectangle                      //구조체 pointer 선선
	rect1 = new(Rectangle)                    //구조체 pointer에 메모리 할당
	rect2 := new(Rectangle)                   //구조체 pointer 선언과 동시에 memory 할당
	var rect1 Rectangle = Rectangle{10, 20}   //구조체 인스턴스 생성하면서 값 초기화
	rect2 := Rectangle{45, 62}                //var keyword 생략. 구조체 인스턴스 생성하면서 값 초기화
	rect3 := Rectangle{width: 30, height: 15} //구조체 필드를 지정해 값 초기화

	var rect1 Rectangle                   //구조체 인스턴스 생성
	var rect2 *Rectangle = new(Rectangle) //구조체 포인터 선언 후 메모리 할당

	rect1.height = 20
	rect2.height = 62

	fmt.Println(rect1) // {0 20}
	fmt.Println(rect2) //&{0 62} --> 구조체 포인터이므로 앞에 & 붙음
}
*/
/*
type Rectangle struct {
	width, height int
}

func NewRectangle(width, height int) *Rectangle {
	return &Rectangle{width, height} //구조체 인스턴스 생성한 뒤 pointer 리턴
}
func main() {
	//rect := NewRectangle(10, 20)
	rect := &Rectangle{20, 10}
	fmt.Println(rect)
}
*/
/*
type Rectangle struct {
	width, height int
}

func rectangleArea(rect *Rectangle) int {
	return rect.width * rect.height
}
func rectangleScaleA(rect *Rectangle, factor int) { //매개변수로 구조체 포인터를 받음
	rect.width = rect.width * factor   //pointer이므로 원래의 값이 변경됨
	rect.height = rect.height * factor //pointer이므로 원래의 값이 변경됨
}
func rectangleScaleB(rect Rectangle, factor int) { //매개변수로 구조체 인스턴스를 받음
	rect.width = rect.width * factor   //값이 복사되었으므로 원래의 값에는 영향을 미치지 않음
	rect.height = rect.height * factor //값이 복사되었으므로 원래의 값에는 영향을 미치지 않음
}
func main() {
	rect1 := Rectangle{30, 30}
	//area := rectangleArea(&rect) //구조체의 pointer를 넘김
	//fmt.Println(area)
	rectangleScaleA(&rect1, 10)
	fmt.Println(rect1)

	rect2 := Rectangle{30, 30}
	rectangleScaleB(rect2, 10)
	fmt.Println(rect2)
}
*/
//goroutine
/*
func hello(n int) {
	r := rand.Intn(100)          //랜덤한 숫자 생성
	time.Sleep(time.Duration(r)) //랜덤한 시간동안 대기ㅣ
	fmt.Println(n)               //n출력
}
func main() {
	for i := 0; i < 100; i++ { //100번 반복해
		go hello(i) //함수를 고루틴으로 실행(고루틴 100개 생)
	}
	fmt.Scanln() //main 함수가 종료되지 않도록 대기
}
*/
//multicore 사용하기
/*
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) //runtime.NumCPU()함수로 CPU 개수를 구한 뒤 runtime.GOMAXPROCS함수에 사용할 최대 CPU 개수 생성
	//runtime.GOMAXPROCS 함수는 CPU 코어 개수를 구하지 않고, 특정 값을 설정해도 됨
	//runtime.GOMAXPROCS 함수에 0을 넣으면 설정 값은 바꾸지 않으며 현재 설정 값만 리턴
	fmt.Println(runtime.GOMAXPROCS(0))

	s := "Hello, world!"

	for i := 0; i < 100; i++ {
		go func(n int) {
			fmt.Println(s, n)
		}(i)
	}
	fmt.Scanln()
}
*/
//closure
/*
func main() {
	runtime.GOMAXPROCS(1) //CPU 하나만 사용
	s := "Hello, world!"
	for i := 0; i < 100; i++ {
		go func(n int) { //익명함수를 고루틴으로 실행(클로저)
			fmt.Println(s, n) //s와 매개변수로 받은 n값 출력
		}(i)
		go func() {
			fmt.Println(s, i)
		}()

	}
	fmt.Scanln()
}
*/

//Channel --> make(chan 자료형)
/*
func sum(a int, b int, c chan int) {
	c <- a + b //channel에 a+b를 보냄
}
func main() {
	//	c := make(chan int) //int형 channel 생성. channel을 사용하기 전에는 반드시 make함수로 공간을 할당해야 함. ==> synchronous channel 생성됨
	var c chan int //chan int형 변수 선언
	c = make(chan int)
	go sum(1, 2, c) //sum을 고루틴으로 실행한 뒤 채널을 매개변수로 넘겨줌

	n := <-c //channel에서 값을 꺼낸 뒤 n에 대입
	//<-c는 채널에서 값이 들어올 떄까지 대기. 이 후 채널에 값이 들어오면 대기를 끝내고 다음 코드 실행 ==> 값을 주고 받는 동시에 동기화 역할까지 수행
	fmt.Println(n)
}
*/
//동기 채널
/*
func main() {
	done := make(chan bool) //동기 채널 생성
	count := 5              //반복할 회시수

	go func() {
		for i := 0; i < count; i++ {
			done <- true                //goroutine에 true 보냄. 값을 꺼낼 떄까지 대기. 따라서 반복문도 실행되지 않으므로 '고루틴:숫자'가 계속해서 출력되지 않음.
			fmt.Println("고루틴 : ", i)    //반복문의 변수 출력
			time.Sleep(1 * time.Second) //1초 대기
		}
	}()
	for i := 0; i < count; i++ {
		<-done                     //channel에 값이 들어올 떄까지 대기, 값을 꺼냄.
		fmt.Println("메인 함수 : ", i) //반복문의 변수 출력
	} //고루틴 -> 메인함수 -> 고루틴 -> 메인함수 순으로 실행됨
}
*/
//채널 버퍼링
//채널에 버퍼를 1개 이상 설정하면 비동기ㅣ 채널이 생성됨.
//비동기 채널은 보내는 쪽에서 버퍼가 가득차면 실행을 멈추고
//대기하며 받는 쪽에서는 버퍼에 값이 없으면 대기한다.
//고루틴 생성후 반복문을 실행할 때마다 채널 done에 true값을 보냄
/*
func main() {
	runtime.GOMAXPROCS(1)
	done := make(chan bool, 2) //buffer가 2개인 비동기 채널 생성
	count := 4                 //반복할 횟수

	go func() {
		for i := 0; i < count; i++ {
			done <- true //channel에 true 보냄. buffer가 가득차면 대기
			//비동기 채널이므로 버퍼가 가득찰떄까지 값을 계속 보냄.
			//여기서는 채널의 버퍼가 2이므로 done에 true를 2번 보낸 뒤 다음 lop에서 대기.
			fmt.Println("고루틴 : ", i)
		}
	}()

	for i := 0; i < count; i++ {
		<-done //buffer에 값이 없으면 대기, 값을 꺼냄.
		//비동기 채널에 버퍼가 2개. 이미 값이 2개 들어있ㅇ음. --> loop을 2번 반복하며 <-done에서 값을 꺼냄.
		//이후 채널이 비었으므로 실행을 멈추고 대기.
		//이후 다시 고루틴쪽에서 값을 두번 보내고, 메인에서 2번 꺼냄.
		// 고루틴 -> 고루틴 -> 메인 -> 메인 순으로 실행
		fmt.Println("메인 함수 : ", i) //반복문의 변수 출력
	}
}
*/
//range, close
//이미 닫힌 채널에 값을 보내면 패닉 발생
//채널을 닫으면 range loop 종료
//채널이 열려있고, 값이 들어오지 않는다면 range는 실행되지 않고 계속 대기.
//만약 다른 곳에서 채널에 값을 보냈다면(채널에 값이 들어오면) 그때부터 range 계속 반복됨
// range는 채널에 값이 몇개나 들어올지 모르기 때문에 값이 들어올 때마다 계속 꺼내기 위해 사용.
/*
func main() {
	c := make(chan int) //int 형 채널 생성
	go func() {
		for i := 0; i < 5; i++ {
			c <- i //채널에 값을 보냄
		}
		close(c)
	}()
	for i := range c { //range 이용해 채널이 닫힐 때까지 반복해서 값을 꺼냄
		fmt.Println(i) //꺼낸 값 출력
	}
}
*/
//mutex
//RWMutext:  읽기/쓰기 mutex. 읽기와 쓰기 동작을 나누어 lock을 걸 수 있다.
//Cond : 조건 변수(condition variable)입니다. 대기하고 있는 하나의 객체를 깨울 수도 있고 여러 개를 동시에 깨울 수도 있습니다.
//Once: 특정 함수를 딱 한 번만 실행할 때 사용합니다.
//Pool: 멀티 스레드(고루틴)에서 사용할 수 있는 객체 풀입니다. 자주 사용하는 객체를 풀에 보관했다가 다시 사용합니다.
//WaitGroup: 고루틴이 모두 끝날 때까지 기다리는 기능입니다.
//Atomic: 원자적 연산이라고도 하며 더 이상 쪼갤 수 없는 연산이라는 뜻입니다. 멀티 스레드(고루틴), 멀티코어 환경에서 안전하게 값을 연산하는 기능입니다.
//sync.Mutex
//func(m *Mutex) Lock() : mutex 잠금
//func(m *Mutex) Unlock() : mutex 잠금 해제
/*
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) //모든 CPU사용
	var data = []int{}                   //int형 slice 생성
	var mutext = new(sync.Mutex)

	go func() { //goroutine에서
		for i := 0; i < 1000; i++ { //1000번 반복하며
			mutext.Lock()          //mutex 잠금, data slice 보호 시작
			data = append(data, 1) //dataslice에 1 추가
			mutext.Unlock()        //mutex 잠금 해제. data slice 보호 종료
			runtime.Gosched()      // 다른 goroutine이 cpu를 사용할 수 있도록 양보
		}
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			mutext.Lock()
			data = append(data, 1)
			runtime.Gosched()
			mutext.Unlock()
		}
	}()
	time.Sleep(2 * time.Second)
	fmt.Println(len(data))

}
*/
//RWMutex
/*
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var data int = 0
	var rwMutex = new(sync.RWMutex) //읽기 쓰기 뮤텍스 생성

	go func() { //값을 쓰는 고루틴
		for i := 0; i < 3; i++ {
			rwMutex.Lock()                //쓰기 뮤텍스 잠금, 쓰기 보호 시작
			data += 1                     // data에 값 쓰기
			fmt.Println("write : ", data) //data값 출력
			time.Sleep(10 * time.Millisecond)
			rwMutex.Unlock() //쓰기 뮤텍스 잠금 해제, 쓰기 보호 종료
		}
	}()

	go func() { //값을 읽는 고루틴
		for i := 0; i < 3; i++ {
			rwMutex.RLock()               //읽기 뮤텍스 잠금, 읽기 보호 시작
			fmt.Println("Read1 : ", data) //data 값을 출력(읽기)
			time.Sleep(1 * time.Second)
			rwMutex.RUnlock() //읽기 뮤텍스 잠금 해제, 읽기 보호 종료
		}
	}()

	go func() { //값을 읽는 고루틴
		for i := 0; i < 3; i++ {
			rwMutex.RLock()               //읽기 뮤텍스 잠금, 읽기 보호 시작
			fmt.Println("Read2 : ", data) //data 값 출력(읽기)
			time.Sleep(2 * time.Second)
			rwMutex.RUnlock() //읽기 뮤텍스 잠금, 읽기 보호 시작
		}
	}()
	time.Sleep(10 * time.Second) //10초동안 프로그램 실행
}
*/
//조건변수 ==> 대기하고 있는 객체 하나만 깨우너가 여러 개를 동시에 깨울 때 사용
//sync.Cond
//func NewCond(l Locker) *Cond: 조건 변수 생성
//func (c *Cond) Wait(): 고루틴 실행을 멈추고 대기
//func (c *Cond) Signal(): 대기하고 있는 고루틴 하나만 깨움
//func (c *Cond) Broadcast(): 대기하고 있는 모든 고루틴을 깨움
/*
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	var mutex = new(sync.Mutex)    // 뮤텍스 생성
	var cond = sync.NewCond(mutex) // 뮤텍스를 이용하여 조건 변수 생성

	c := make(chan bool, 3) // 비동기 채널 생성

	for i := 0; i < 3; i++ {
		go func(n int) { // 고루틴 3개 생성
			mutex.Lock() // 뮤텍스 잠금, cond.Wait() 보호 시작
			c <- true    // 채널 c에 true를 보냄
			fmt.Println("wait begin : ", n)
			cond.Wait() // 조건 변수 대기
			fmt.Println("wait end : ", n)
			mutex.Unlock() // 뮤텍스 잠금 해제, cond.Wait() 보호 종료
		}(i)
	}

	for i := 0; i < 3; i++ {
		<-c // 채널에서 값을 꺼냄, 고루틴 3개가 모두 실행될 때까지 기다림
	}

	for i := 0; i < 3; i++ {
		mutex.Lock() // 뮤텍스 잠금, cond.Signal() 보호 시작
		fmt.Println("signal : ", i)
		cond.Signal()  // 대기하고 있는 고루틴을 하나씩 깨움
		mutex.Unlock() // 뮤텍스 잠금 해제, cond.Signal() 보고 종료
	}

	fmt.Scanln()
}
*/
/*
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	var mutex = new(sync.Mutex)    // 뮤텍스 생성
	var cond = sync.NewCond(mutex) // 뮤텍스를 이용하여 조건 변수 생성

	c := make(chan bool, 3) // 비동기 채널 생성

	for i := 0; i < 3; i++ {
		go func(n int) { // 고루틴 3개 생성
			mutex.Lock() // 뮤텍스 잠금, cond.Wait() 보호 시작
			c <- true    // 채널 c에 true를 보냄
			fmt.Println("wait begin : ", n)
			cond.Wait() // 조건 변수 대기
			fmt.Println("wait end : ", n)
			mutex.Unlock() // 뮤텍스 잠금 해제, cond.Wait() 보호 종료

		}(i)
	}

	for i := 0; i < 3; i++ {
		<-c // 채널에서 값을 꺼냄, 고루틴 3개가 모두 실행될 때까지 기다림
	}

	mutex.Lock() // 뮤텍스 잠금, cond.Broadcast() 보호 시작
	fmt.Println("broadcast")
	cond.Broadcast() // 대기하고 있는 모든 고루틴을 깨움
	mutex.Unlock()   // 뮤텍스 잠금 해제, cond.Signal() 보고 종료

	fmt.Scanln()
}
*/
/*
import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {

	// Must Kept Secret No Hardcoding , This is for Demo purpose.
	key := "myverystrongpasswordo32bitlength"

	// IN CBC Must be Block Size of AES (Multiple of 16)
	// Other WIse Paddign needs to be perfomed
	plainText := "1234567890123456"

	if len(plainText)%aes.BlockSize != 0 {
		panic("plaintext is not a multiple of the block size")
	}

	fmt.Printf("Original Text:  %s\n", plainText)
	fmt.Println("====CBC Encryption/ Decryption====")
	// IV Length Must be equal to Block Size.
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err.Error())
	}
	ciphertext := CBCEncrypter(key, plainText, iv, nil)
	fmt.Printf("CBC Encrypted Text:  %s\n", ciphertext)
	ret := CBCDecrypter(key, ciphertext, iv, nil)
	fmt.Printf("CBC Decrypted Text:  %s\n", ret)
}

func CBCEncrypter(key string, plaintext string, iv []byte, additionalData []byte) string {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], []byte(plaintext))
	return hex.EncodeToString(ciphertext)
}

func CBCDecrypter(key string, ct string, iv []byte, additionalData []byte) string {
	ciphertext, _ := hex.DecodeString(ct)
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	ciphertext = ciphertext[aes.BlockSize:]
	// CBC mode always works in whole blocks.
	if len(ciphertext)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(ciphertext, ciphertext)
	s := string(ciphertext[:])

	return s
}
*/ /*
import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(3000 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}

	}()
	time.Sleep(160000 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker Stopped")
}
*/
/*
import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Minute)
	fmt.Println(time.Now().Unix())
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(10 * time.Minute)
		done <- true //채널에 'true' 보냄
	}()
	for {
		select {
		case <-done: //
			fmt.Println("Done!")
			return
		case t := <-ticker.C: //채널로부터 값을 받으면 t에 입력
			fmt.Println("Current time : ", t)
		}
	}
}
*/
import (
	"fmt"
	"os"
	"strings"
)

func main() {
	tables := []string{"policy_list", "traffic_log", "detect_log", "id_manage", "server_status_log"}
	for _, e := range tables {
		os.Setenv("table", e)
	}
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0], "=>", pair[1])
	}
}
