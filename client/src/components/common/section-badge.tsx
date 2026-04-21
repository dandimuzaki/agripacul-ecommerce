import React from 'react'

const SectionBadge = ({text}: {text: string}) => {
  return (
    <div className="mb-4 text-primary-dark bg-primary-dark/20 px-3 py-1 rounded-full font-medium w-fit uppercase text-xs md:text-sm">{text}</div>
  )
}

export default SectionBadge
