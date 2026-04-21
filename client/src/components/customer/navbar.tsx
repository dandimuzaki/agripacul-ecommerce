'use client';

import React, { useEffect, useState } from 'react'
import { InputGroup, InputGroupInput } from '@/components/ui/input-group'
import { Menu, SearchIcon, X } from 'lucide-react'
import Image from 'next/image'
import { usePathname, useRouter, useSearchParams } from 'next/navigation';
import Link from 'next/link';
import { Button } from '../ui/button';
import CartButton from '../common/cart-button';
import { LocalMallRounded } from '@mui/icons-material';
import { useProfile } from '@/hooks/profile/useProfile';
import { CustomerMenu } from './customer-menu';
import { AnimatePresence, motion } from 'framer-motion';

const Navbar = () => {
  const [scrolled, setScrolled] = useState(false);
  const [keyword, setKeyword] = useState("");
  const searchParams = useSearchParams();
  const router = useRouter()

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

  const onSubmit = (keyword: string) => {
    const params = new URLSearchParams(searchParams.toString())    
    params.set("search", keyword)
    router.push(`/products?${params.toString()}`)
  }

  const { data: profile } = useProfile()

  const [isOpen, setIsOpen] = useState(false);

  const menuVariants = {
    hidden: { opacity: 0, y: -20 },
    visible: {
      opacity: 1,
      y: 0,
      transition: {
        staggerChildren: 0.1,
        duration: 0.3,
      },
    },
    exit: { opacity: 0, y: -20, transition: { duration: 0.2 } },
  };

  const itemVariants = {
    hidden: { opacity: 0, y: -10 },
    visible: { opacity: 1, y: 0 },
  };

  const menu = [
    {
      href: "/products",
      title: "Browse Products",
      mustAuth: false
    },
    {
      href: "/orders",
      title: "My Orders",
      mustAuth: true
    },
    {
      href: "/profile",
      title: "My Profile",
      mustAuth: true
    },
    {
      href: "/contact",
      title: "Contact Us",
      mustAuth: false
    },
  ]

  return (
    <nav className={`fixed w-full z-50 transition-all h-12 md:h-16 px-2 py-1 md:px-8 md:py-2 flex justify-between md:grid md:grid-cols-[2fr_3fr_2fr] items-center gap-2 ${!isHome ? 'bg-primary' : isOpen ? 'bg-primary' : scrolled ? 'bg-primary' : 'bg-transparent'}`}>
      {/* Logo */}
      <div className='w-36 md:w-fit'>
        <div className='md:h-8 w-full'>
          <Link href="/" className='w-full'>
            <Image src={!isHome ? "/logo.png" : isOpen ? "/logo.png" : scrolled ? "/logo.png" : "/logo-colored.png"} height={100} width={100} alt='logo Agripacul' className={`${!isHome ? 'md:h-8' : scrolled ? 'md:h-8' : 'md:h-12'} object-cover h-full w-full`} />
          </Link>
        </div>
      </div>

      {/* Desktop */}
      <InputGroup className={`${!isHome ? 'max-w-lg' : scrolled ? 'max-w-lg' : 'w-0'} hidden md:flex rounded-full bg-white border-0 overflow-hidden`}>
        <InputGroupInput
          className='h-4'
          placeholder="Search products..."
          value={keyword}
          onChange={(e) => setKeyword(e.target.value)}
        />
        <Button variant="searchIcon" type='button' className='bg-white' onClick={() => onSubmit(keyword)}>
          <SearchIcon/>
        </Button>
      </InputGroup>
      <div className='hidden md:flex justify-end gap-2 items-center w-full'>
        {profile && <div className='flex items-center gap-2'>
          <CartButton isHome={isHome} isOpen={isOpen} scrolled={scrolled}/>
        </div>}
          {profile && <>
            <CustomerMenu color={isHome && !scrolled} profile={profile}/>
          </>}
          {!profile && <>
          <Link href="/login">
            <Button variant="outline" className={`${!isHome ? 'border-white text-white' : scrolled ? 'border-white text-white' : ''} bg-transparent hover:bg-primary-foreground hover:border-primary-foreground border `}>Login</Button>
          </Link>
          <Link href="/register">
            <Button className={`${!isHome ? 'bg-white text-primary hover:text-white' : scrolled ? 'bg-white text-primary hover:text-white' : 'bg-primary text-white'} hover:bg-primary-foreground`}>Sign Up</Button>
          </Link>
          </>}
      </div>

      {/* Mobile Button */}
      <div className='md:hidden flex gap-2 items-center'>
        <div className={profile ? '' : 'hidden'}><CartButton isHome={isHome} isOpen={isOpen} scrolled={scrolled}/></div>
        <button
          className="md:hidden"
          onClick={() => setIsOpen(!isOpen)}
        >
          {isOpen ? <X /> : <Menu />}
        </button>
      </div>

      {/* Animated Mobile Menu */}
      <AnimatePresence>
        {isOpen && (
          <motion.div
            variants={menuVariants}
            initial="hidden"
            animate="visible"
            exit="exit"
            className="absolute top-12 left-0 w-full bg-primary backdrop-blur-md flex flex-col items-center py-4 md:hidden shadow-lg"
          >
            {menu.filter(item => profile ? item : !item.mustAuth).map((item) => (
              <motion.div key={item.title} variants={itemVariants} className='w-full h-full px-4 py-2 hover:bg-primary-foreground'>
                <Link
                  href={`#${item.href}`}
                  onClick={() => setIsOpen(false)}
                  className="text-white w-full flex"
                >
                  {item.title}
                </Link>
              </motion.div>
            ))}

            {!profile && <motion.div variants={itemVariants} className="flex gap-2 mt-4">
              
                <>
                  <Link href="/login">
                    <Button variant="outline" className={`border-white text-white bg-transparent hover:bg-primary-foreground hover:border-primary-foreground border `}>Login</Button>
                  </Link>
                  <Link href="/register">
                    <Button className={`bg-white text-primary hover:text-white bg-white text-primary hover:text-white' : 'bg-primary text-white'} hover:bg-primary-foreground`}>Sign Up</Button>
                  </Link>
                </>
              
            </motion.div>}
          </motion.div>
        )}
      </AnimatePresence>
    </nav>
  );
}

export default Navbar
