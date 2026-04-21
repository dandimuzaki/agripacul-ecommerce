"use client"
import { Audiotrack, Email, Instagram, LinkedIn, LocationOn, Twitter, WhatsApp, X } from '@mui/icons-material';
import { Button } from '../ui/button';
import Link from 'next/link';

const Footer = () => {
  return (
    <div className='bg-primary-dark overflow-hidden h-12 md:h-full'>
      <div className='flex justify-center md:grid gap-4 px-8 lg:px-16 py-4 lg:py-8 text-white md:grid-cols-[2fr_1fr_1fr]'>
      <div className='space-y-4'>
        <div className='flex justify-center items-center w-28 md:w-48'><img src={"/logo.png"} className="w-full h-full" /></div>
        <p className='hidden md:block text-sm/5'>Agripacul connects you with fresh produce, homemade foods, and gardening essentials, grown responsibly and delivered with care for everyday cooking and growing</p>
        <div className='hidden md:flex justify-center md:justify-start gap-2'>
          <Button className="aspect-square rounded-full h-full" ><Instagram fontSize='small'/></Button>
          <Button className="aspect-square rounded-full h-full"><X fontSize='small'/></Button>
          <Button className="aspect-square rounded-full h-full"><Audiotrack fontSize='small'/></Button>
          <Button className="aspect-square rounded-full h-full"><LinkedIn fontSize='small'/></Button>
        </div>
      </div>
      <div className='hidden md:block space-y-2'>
        <p className='font-bold text-primary'>Quick Links</p>
        <div className="text-sm flex flex-col gap-1">
          <Link href={"/"} className='hover:text-primary'>Home</Link>
          <Link href={"/products"} className='hover:text-primary'>Browse Products</Link>
          <Link href={"/orders"} className='hover:text-primary'>Track Your Order</Link>
          <Link href={"/about"} className='hover:text-primary'>About Us</Link>
        </div>
      </div>
      <div className='hidden md:block space-y-2'>
        <p className='text-primary font-bold'>Get in Touch</p>
        <a className='flex gap-2 text-sm cursor-pointer items-center'><WhatsApp />+62 853-2409-1088</a>
        <a className='flex gap-2 text-sm cursor-pointer items-center'><Email />agripacul@itb.lpik.org</a>
        <a className='flex gap-2 text-sm cursor-pointer'><LocationOn/>Jl. Cisintok Kadumulya, Cihanjuang, Kec. Parongpong, Kabupaten Bandung Barat</a>
      </div>
      </div>
    </div>
  );
};

export default Footer;