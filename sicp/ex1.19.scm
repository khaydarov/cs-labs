;Consider T to be transformation of a pair (a, b).
;Special case of Tpq is the transformation with the rules: a = bq+aq+ap, b = bp+qp
;If we apply Tpq twice it means: 
; - Tpq(a,b) = (bq+aq+ap, bp+aq) - first time
; - Tpq(Tpq(a,b)) = ((bp+aq)q + (bq+aq+ap)q + (bq+aq+ap)p, (bp+aq)p + (bq+aq+ap)q) - second time
; So, we need to prove that Tp'q' is equals to Tpq(Tpq(a,b)) and compute p', q' in terms of p, q
; Let's make equation of two this expressions:
; Tp'q' = (bq'+aq'+ap', bp'+aq')
; bq'+aq'+ap' = (bp+aq)q + (bq+aq+ap)q + (bq+aq+ap)p
; Now presume a pair (a, b) as (1, 0) (given in exersice). Then we get q'+p'=q^2+(p+q)^2
; bp'+aq' = (bp+aq)p + (bq+aq+ap)q, - q' = q^2 + 2qp
; Finally,
; p' = p^2 + q^2
; q' = q^2 + 2qp
(define (fib n)
  (define (fib-iter a b p q count)
	(cond ((= count 0) b)
			((even? count) (fib-iter 
									a
									b
									(+ (square p) (square q))
									(+ (square q) (* 2 q p))
									(/ count 2)))
			(else (fib-iter
							(+ (* b q) (* a q) (* a p))
							(+ (* b p) (* a q))
							p
							q
							(- count 1)))))
  (fib-iter 1 0 0 1 n))