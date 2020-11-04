;Returns 10
10

;Returns 12
(+ 5 3 4)

;Returns 8
(- 9 1)

;Returns 3
(/ 6 2)

;Returns 6
(+ (* 2 4) (- 4 6))

;Sets a to 3
(define a 3)

;Sets b to 4
(define b (+ a 1))

;Returns 19
(+ a b (* a b))

;Returns false
(= a b)

;Returns b: 4
(if 
  (and (> b a) (< b (* a b)))
	b
	a)

;Returns 16
(cond 
  ((= a 4) 6)
  ((= b 4) (+ 6 7 a))
  (else 25))

;Returns 6
(+ 2 (if (> b a) b a))

;Returns 16
(* (cond ((> a b) a)
     	((< a b) b)
     	(else -1))
  	(+ a 1))