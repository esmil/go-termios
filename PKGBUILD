# Contributor: Esmil <esmil@mailme.dk>

pkgname=go-termios
pkgver=$(date +%Y%m%d)
pkgrel=1
pkgdesc='Go code to get and set termios'
arch=('i686' 'x86_64')
url='http://github.com/esmil/go-termios'
license=('GPL')
depends=('go>=2010_08_25')
options=(!strip)

build() {
  cd ${srcdir}
  ln -s .. $pkgname
  cd $pkgname

  make
}

package() {
  cd ${srcdir}/$pkgname

  make DESTDIR=${pkgdir} install
}
