;Simple combination: sum two operands
(+ 1 3)

;Compound (nested) combination: sum of two expressions
(+ (+ 1 3) (* 2 4))

;Complex arithmetic expression
(+
  (* (+ 1 5) (+ 1 6))
  (* (* 2 5) (+ 3 4))
)

;Sets variable value
(define foo 4)

;Defines a "function" that squares the given value
(define (square x) (* x x))

;Defines a new function that uses another function
;Lisp uses applicative-order evaluation
(define
  (sum_of_squares a b)
  	(+ (square a) (square b)))

;Case analysis
;Function: abs(x) -> if x > 0: x; if x = 0: 0; if x < 0: -x;
;Syntax:
;  (cond (<p1> <e1>)
;  		 (<p2> <e2>)
;  		 (<p3> <e3>)
;  		 ...
;  		 (<pN> <eN>)
(define (abs x)
  (cond ((> x 0) x)
    	((= x 0) 0)
    	((< x 0) (- x))))

;Alternative way of case analysis
(define (abs1 x)
  (cond ((< x 0) (- x))
    	(else x)))

;Case analysis with if
;Syntax:
;	(if <predicate> <consequent> <alternative>)
(define (abs2 x)
  (if (< x 0)
     (- x)
     x))

;Custom predicates
(define (>= x y) (or (> x y) (= x y)))
(define (number_between x y z) (and (< x z) (> y z)))

;Tip:
; - functions defined with lambda use applicative order and evaluated right to left
; - special forms use normal order (and, or, if, cond ..)
;
; Applicative-order
; (double (* 3 4))
; (double 12)
; (+ 12 12)
; 24
;
; Normal-order
; (double (* 3 4))
; (+ (* 3 4) (* 3 4))
; (+ 12 (* 3 4))
; (+ 12 12)
; 24

;Recursion: calculate factorial
define (factorial n)
  	(if (= n 1)
  		1
  		(* n (factorial (- n 1)))))

;Closures
;Lets calculate sum: 1 + 2 + 3 + ... + n
;For that we create inc procedure
(define (inc a) (+ a 1))
;then identity (returns argument itself)
(define (identity x) x)
;And write sum procedure
(define (sum term a next b)
	(if (> a b)
		0
		(+ (term a)
			(sum term (next a) next b))))
;So sum of numbers will be
(define (sum-int a b)
  (sum identity a inc b))

;Or sum of squares
(define (sum-sq a b)
  (sum square a inc b))

;Lambda procedure: anonymous functions
;Returns procedure that can be used
(lambda (a b) (+ a b))

;Local variables
(define (f)
  (let ((a 1)
    (b 2))
  (+ a b)))

;defining list
(cons 1
    (cons 2
        (cons 3
            (cons 4 nil))))

;or use syntax list
(list 1 2 3 4)
