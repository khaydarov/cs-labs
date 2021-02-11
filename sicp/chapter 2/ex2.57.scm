(define (variable? x) (symbol? x))
(define (=number? x v) (and (number? x) (= x v)))
(define (same-variable? v1 v2)
	(and (variable? v1) (variable? v2) (eq? v1 v2)))

(define (sum? x) (and (pair? x) (eq? (cadr x) '+)))
(define (addend x) (car x))
(define (augend x) (caddr x))
(define (make-sum a1 a2)
	(cond ((=number? a1 0) a2)
				((=number? a2 0) a1)
				((and (number? a1) (number? a2)) (+ a1 a2))
				(else (list a1 '+ a2))))

(define (product? x) (and (pair? x) (eq? (cadr x) '*)))
(define (multiplier x) (car x))
(define (multiplicand x) (caddr x))
(define (make-product p1 p2)
	(cond ((or (=number? p1 0) (=number? p2 0)) 0)
				((=number? p1 1) p2)
				((=number? p2 1) p1)
				((and (number? p1) (number? p2)) (* p1 p2))
				(else (list p1 '* p2))))

(define (exponentiation? x) (and (pair? x) (eq? (cadr x) '**)))
(define (base x) (car x))
(define (exponent x) (caddr x))
(define (make-exponentiation e1 e2)
	(cond ((=number? e2 0) 1)
			  ((=number? e2 1) e1)
				((and (number? e1) (number? e2)) (expt e1 e2))
				(else (list e1 '** e2))))

(define (deriv exp var)
	(cond ((number? exp) 0)
			  ((variable? exp) (if (same-variable? exp var) 1 0))
				((sum? exp)
					(make-sum (deriv (addend exp) var)
										(deriv (augend exp) var)))
				((product? exp)
					(make-sum (make-product (multiplier exp)
																	(deriv (multiplicand exp) var))
										(make-product (deriv (multiplier exp) var)
																	(multiplicand exp))))
				((exponentiation? exp)
					(make-product
						(make-product (exponent exp)
													(make-exponentiation (base exp) (- (exponent exp) 1)))
						(deriv (base exp) var)))
				(else
					(error "unknown expression type"))))
