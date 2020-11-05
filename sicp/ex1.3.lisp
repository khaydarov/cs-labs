(define (square x) (* x x))
(define (square_sum a b) (+ (square a) (square b)))
(define (my_func a b c)
	(cond 
		((and (> a b) (> b c)) (square_sum a b))
    	((and (> a b) (> c b)) (square_sum a c))
		((and (> b a) (> c b)) (square_sum b c))))