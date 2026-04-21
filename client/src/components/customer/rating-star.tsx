"use client";

import { Star, StarBorder, StarHalf } from "@mui/icons-material";

type RatingStarsProps = {
  rating: number; // ex: 4.5
  max?: number;   // default 5
};

export const RatingStars = ({ rating, max = 5 }: RatingStarsProps) => {
  const fullStars = Math.floor(rating);
  const hasHalfStar = rating - fullStars >= 0.5;
  const emptyStars = max - fullStars - (hasHalfStar ? 1 : 0);

  return (
    <div className="flex items-center text-yellow-500">
      {/* Full stars */}
      {Array.from({ length: fullStars }).map((_, i) => (
        <Star fontSize="small" key={i} />
      ))}

      {/* Half star */}
      {hasHalfStar && (
          <StarHalf fontSize="small"/>
      )}

      {/* Empty stars */}
      {Array.from({ length: emptyStars }).map((_, i) => (
        <StarBorder fontSize="small" key={i}/>
      ))}
    </div>
  );
};