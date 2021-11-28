=====================================================
Correct Timestamp of Google Photo from Google Takeout
=====================================================

.. image:: https://img.shields.io/badge/Language-Go-blue.svg
   :target: https://golang.org/

.. image:: https://godoc.org/github.com/siongui/go-fix-google-photo-timestamp?status.svg
   :target: https://godoc.org/github.com/siongui/go-fix-google-photo-timestamp

.. image:: https://github.com/siongui/go-fix-google-photo-timestamp/workflows/ci/badge.svg
    :target: https://github.com/siongui/go-fix-google-photo-timestamp/blob/master/.github/workflows/ci.yml

.. image:: https://goreportcard.com/badge/github.com/siongui/go-fix-google-photo-timestamp
   :target: https://goreportcard.com/report/github.com/siongui/go-fix-google-photo-timestamp

.. image:: https://img.shields.io/badge/license-Unlicense-blue.svg
   :target: https://github.com/siongui/go-fix-google-photo-timestamp/blob/master/UNLICENSE


Development Environment:

  - `Ubuntu 20.04`_
  - `Go 1.17`_


Usage
+++++

Require Go 1.16 or later to build.

.. code-block:: bash

  go build -o fixtime

.. code-block:: bash

  ./fixtime -dir=/path/to/your/folder


UNLICENSE
+++++++++

Released in public domain. See UNLICENSE_.


References
++++++++++

.. [1] `Coming in Go 1.16: ReadDir and DirEntry <https://benhoyt.com/writings/go-readdir/>`_


.. _Go: https://golang.org/
.. _Ubuntu 20.04: https://releases.ubuntu.com/20.04/
.. _Go 1.17: https://golang.org/dl/
.. _UNLICENSE: https://unlicense.org/
.. _os.ReadDir: https://pkg.go.dev/os#ReadDir
