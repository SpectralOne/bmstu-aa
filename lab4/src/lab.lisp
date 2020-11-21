(ql:quickload :array-operations)
(ql:quickload :lparallel)

(setf lparallel:*kernel* 
    (lparallel:make-kernel 8))

(defun init-matrix (rows columns) 
    (aops:generate (lambda () (random 10)) (list rows columns)))

(defun mult-row (m1-row-idx m2-col-idx)
    (setq sum 0)
    (loop for i from 0 below (length (aops:sub matrix1 m1-row-idx))
        do (setf sum  
            (+ sum (* (aref (aops:sub matrix1 m1-row-idx) i) 
                      (aref (aops:sub (list-to-2d-array 
                                      (rotate (2d-array-to-list matrix2))) 
                                       m2-col-idx) i)))))

    (setf (aref new-matrix m1-row-idx m2-col-idx) sum))

(defun mult-matrix (matrix1 matrix2)
    (destructuring-bind (n m) (array-dimensions matrix1)
        (setq m1-rows n)
        (setq m1-cols m))

    (destructuring-bind (n m) (array-dimensions matrix2)
        (setq m2-cols m)
        (setq m2-rows n))

    (defvar new-matrix-rows m1-rows)
    (defvar new-matrix-columns m2-cols)

    (if (/= m2-rows m1-cols)
        (return-from mult-matrix 
            (PRINT "Wrong dimensions!")))

    (defvar new-matrix (aops:generate 
        (lambda () ()) 
            (list new-matrix-rows new-matrix-columns)))

    (time (lparallel:pdotimes (i m1-rows)
        (dotimes (j m2-cols)
            (mult-row i j))))
    (PRINT new-matrix))

(defun rotate (list-of-lists)
  (apply #'mapcar #'list list-of-lists))

(defun 2d-array-to-list (array)
  (loop for i below (array-dimension array 0)
        collect (loop for j below (array-dimension array 1)
                      collect (aref array i j))))

(defun list-to-2d-array (list)
  (make-array (list (length list)
                    (length (first list)))
              :initial-contents list))


(defvar n 3)

(defvar matrix1 (init-matrix n n))
(defvar matrix2 (init-matrix n n))
(defvar cnt 0)

(princ "Size: ")
(write n)
(princ " X ")
(write n)

(PRINT matrix1)
(PRINT matrix2)
;; (dotimes (j 10)
(mult-matrix matrix1 matrix2)

(EXIT)
