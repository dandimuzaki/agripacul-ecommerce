import { Sidebar } from 'lucide-react';
import React from 'react'

export default function Admin1Layout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <div>
      <Sidebar/>
      <div className='ml-100'>
        {children}
      </div>
    </div>
  );
}

