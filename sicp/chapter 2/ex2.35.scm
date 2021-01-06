(define nil '())
(define (accumulate op initial sequence)
	(if (null? sequence)
		initial
		(op (car sequence)
			(accumulate op initial (cdr sequence)))))

(define (count-leaves-recursive t)
   (accumulate + 0
     (map
       (lambda (t)
         (cond ((null? t) 0)
               ((pair? t) (count-leaves-recursive t))
               (else 1)))
       t)))
