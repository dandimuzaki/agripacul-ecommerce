'use client';

import { useEffect, useState } from 'react'
import Image from 'next/image'
import { usePathname } from 'next/navigation';
import Link from 'next/link';

const StrictNavbar = () => {
  const [scrolled, setScrolled] = useState(false);

  useEffect(() => {
    const handleScroll = () => {
      const offset = window.scrollY;
      setScrolled(offset > 50);
    };

    window.addEventListener('scroll', handleScroll);
    return () => window.removeEventListener('scroll', handleScroll);
  }, []);

  const pathname = usePathname()
  const isHome = pathname === '/'

  return (
    <div className={`transition-all fixed w-full z-20 ${!isHome ? '' : scrolled ? '' : ''}`}>
      <div className={`transition-all h-12 md:h-16 px-2 py-1 md:px-4 md:py-2 flex items-center gap-2 ${!isHome ? 'bg-primary' : scrolled ? 'bg-primary' : 'bg-transparent'}`}>
        <div className='w-36 md:w-fit'>
          <div className='md:h-8 w-full'>
            <Link href="/" className='w-full'>
              <Image src={"/logo.png"} height={100} width={100} alt='logo Agripacul' className={`${!isHome ? 'md:h-8' : scrolled ? 'md:h-8' : 'md:h-12'} object-cover h-full w-full`} />
            </Link>
          </div>
        </div>
      </div>
    </div>
  )
}

export default StrictNavbar
