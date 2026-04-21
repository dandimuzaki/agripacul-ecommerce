import Contact from "@/components/common/contact-us";
import SectionBadge from "@/components/common/section-badge";
import StorySection from "@/components/common/story-section";
import TestimonialCarousel from "@/components/common/testimonials";
import ValuePreposition from "@/components/common/value-preposition";
import BestSelling from "@/components/customer/best-selling";
import CategoryFilterDropdown from "@/components/customer/category-dropdown";
import { Button } from "@/components/ui/button";
import { InputGroup, InputGroupInput } from "@/components/ui/input-group";
import { SearchIcon } from "lucide-react";
import Link from "next/link";
import InfiniteSlider from "./infinite-slider";
import Tags from "./tags";

export default function Home() {
  return (
      <main className="min-h-screen">
        <section className="grid lg:grid-cols-2 gap-2 lg:gap-6 lg:h-screen">

          <div className="w-full order-2 lg:order-1 text-center lg:text-left gap-4 z-3 relative h-full w-full px-4 lg:px-16 flex flex-col justify-center items-center lg:items-start py-4 lg:py-16">
              <h1 className="text-3xl md:text-5xl text-primary font-bold uppercase">
                From our farms <br/>To your table 
              </h1>
              <span className="text-base/5 md:text-lg/6">
                Discover responsibly grown produce and gardening essentials. Everything you need to cook well and grow your own.
              </span>
              <InputGroup className={`p-2 rounded-full bg-white border-0 overflow-hidden`}>
                <CategoryFilterDropdown/>
                <InputGroupInput
                  placeholder="Search products..."
                  // value={keyword}
                  // onChange={(e) => setKeyword(e.target.value)}
                />
                <Button variant="searchIcon" type='button'>
                  <SearchIcon/>
                </Button>
              </InputGroup>
              <Tags/>
              <div className="flex gap-2 md:gap-4 text-white items-center justify-center md:justify-start">
                <Link href={"/products"}>
                  <Button className="md:text-base font-semibold md:py-5 md:px-5">Shop Now</Button>
                </Link>
                <Link href={"/about-us"}>
                  <Button className="md:text-base border-3 font-semibold md:py-4 md:px-4" variant={"outline"}>Book a Visit</Button>
                </Link>
              </div>
          </div>

          <div className="order-1 lg:order-2 pt-16 px-4 lg:py-20 lg:pr-16 h-72 lg:h-full">
            <InfiniteSlider/>
          </div>
          
        </section>
        <section className="p-4 md:p-16 flex flex-col items-center bg-white">
          <SectionBadge text="What We Provide"/>
          <h2 className="mb-3 md:mb-5 text-2xl/7 md:text-4xl text-primary font-semibold text-center">Everything You Need<br/>to Eat Fresh & Grow More</h2>
          <div className="grid md:grid-cols-2 lg:grid-cols-[3fr_3fr_4fr] lg:grid-rows-8 lg:gap-4 gap-2 lg:items-end lg:h-160 w-full">
            <div className="rounded-lg overflow-hidden h-48 lg:h-full w-full relative lg:col-span-2 lg:row-span-5">
              <img src="/harvest.png" alt="Fresh Vegetables by Agripacul" width={100} height={100} className="h-full w-full object-cover relative" />
              <div className="p-6 flex flex-col gap-1 justify-end absolute z-2 bottom-0 h-full w-full bg-[linear-gradient(rgba(0,0,0,0),rgba(0,0,0,0.8))]">
                <p className="text-xl text-white font-semibold">Fresh Vegetables</p>
                <p className="text-md/4 text-white">Harvested from our 2020m² farm, delivered with care.</p>
              </div>
            </div>
            <div className="rounded-lg overflow-hidden relative h-48 lg:h-full lg:row-span-4">
              <img 
              src="/salad-2.jpg" 
              alt="Salad Shake by Agripacul" 
              width={100} height={100} 
              className="h-full w-full object-cover"
              loading="eager"
               />
              <div className="p-6 flex flex-col gap-1 justify-end absolute z-2 bottom-0 h-48 lg:h-full w-full bg-[linear-gradient(rgba(0,0,0,0),rgba(0,0,0,0.8))]">
                <p className="text-xl text-white font-semibold">Ready-to-Shake Salads</p>
                <p className="text-md/4 text-white">Japanese & Western styles. Just shake, and enjoy the perfect bite.</p>
              </div>
            </div>
            <div className="rounded-lg overflow-hidden h-48 lg:h-full relative lg:row-span-4">
              <img src="/kimchi.jpg" alt="Cherry Tomato by Agripacul" width={100} height={100} className="h-full w-full object-cover" />
              <div className="p-6 flex flex-col gap-1 justify-end absolute z-2 bottom-0 h-full w-full bg-[linear-gradient(rgba(0,0,0,0),rgba(0,0,0,0.8))]">
                <p className="text-xl text-white font-semibold">Homemade & Specialty Foods</p>
                <p className="text-md/4 text-white">Including our signature kimchi and fresh ingredients.</p>
              </div>
            </div>
            <div className="rounded-lg overflow-hidden h-48 lg:h-full relative lg:row-span-3">
              <img src="/seeds.jpg" alt="Cherry Tomato by Agripacul" width={100} height={100} className="h-full w-full object-cover" />
              <div className="p-6 flex flex-col gap-1 justify-end absolute z-2 bottom-0 h-full w-full bg-[linear-gradient(rgba(0,0,0,0),rgba(0,0,0,0.8))]">
                <p className="text-xl text-white font-semibold">Seeds & Seedlings</p>
                <p className="text-md/4 text-white">Start your own mini farm at home.</p>
              </div>
            </div>

            <div className="rounded-lg overflow-hidden h-48 lg:h-full relative lg:row-span-3">
              <img src="/gardening-tools.jpeg" alt="Cherry Tomato by Agripacul" width={100} height={100} className="h-full w-full object-cover" />
              <div className="p-6 flex flex-col gap-1 justify-end absolute z-2 bottom-0 h-full w-full bg-[linear-gradient(rgba(0,0,0,0),rgba(0,0,0,0.8))]">
                <p className="text-xl text-white font-semibold">Gardening Tools</p>
                <p className="text-md/4 text-white">Fun gardening with handy tools.</p>
              </div>
            </div>
            </div>
        </section>
        <section className="p-4 md:p-16">
          <h2 className="mb-3 md:mb-5 text-2xl/7 md:text-4xl text-primary font-semibold">Best Seller</h2>
          <BestSelling/>
        </section>
        <section className="p-4 md:p-16 flex flex-col items-center bg-primary-dark">
          <StorySection/>
        </section>
        <section className="p-4 md:p-16 flex flex-col items-center bg-white">
          <SectionBadge text="Our Value"/>
          <h2 className="mb-3 md:mb-5 text-2xl/7 md:text-4xl text-primary font-semibold text-center">Why People Choose Agripacul</h2>
          <div className="grid md:grid-cols-2 lg:grid-cols-[5fr_3fr_3fr] gap-y-4 gap-x-6">
            <div className="hidden lg:block row-span-2 gap-2">
              <div className="rounded-lg overflow-hidden h-full relative ">
                <img src="/testimonial3.jpg" alt="Fresh Vegetables by Agripacul" width={100} height={100} className="h-full w-full object-cover relative" />
                <div className="p-6 flex flex-col gap-1 justify-end absolute z-2 bottom-0 h-full w-full bg-[linear-gradient(rgba(0,0,0,0),rgba(0,0,0,0),rgba(0,0,0,0.8))]">
                </div>
              </div>
            </div>
            <ValuePreposition/>
          </div>
        </section>
        <section className="p-4 md:p-16 flex flex-col items-center">
          <SectionBadge text="Testimonials"/>
          <h2 className="mb-3 md:mb-5 text-2xl/7 md:text-4xl text-primary font-semibold text-center">What Our Customers Say</h2>
          <TestimonialCarousel/>
        </section>
        <section className="relative">
          <div className="top-0 z-0 absolute h-full w-full">
            <img className="absolute top-0 h-full w-full object-cover" src="/hero.jpg" alt="Vegetable Garden" height={100} width={100}/>
            <div className="absolute bottom-0 h-full w-full bg-[linear-gradient(rgba(17,87,65,0.5))]"></div>
          </div>
          <div className="relative h-full w-full p-4 md:p-16 flex flex-col items-center">
            <h2 className="text-primary text-2xl/7 md:text-4xl font-semibold mb-2 text-center">Let's keep agriculture sustainable!</h2>
            <p className="text-white mb-2 md:mb-6 text-center">Explore our products and start your journey with Agripacul today.</p>
            <div className="flex gap-4 text-white items-center">
                <Link href={"/products"}>
                  <Button className="text-lg font-semibold py-3 px-4">Shop Now</Button>
                </Link>
                <Link href={"/about-us"}>
                  <Button className="text-lg border-2 font-semibold py-3 px-4" variant={"outline"}>Visit Our Farm</Button>
                </Link>
              </div>
          </div>
        </section>
        <section className="p-4 md:p-16 hidden md:flex flex-col items-center">
          <SectionBadge text="Contact Us"/>
          <h2 className="mb-3 md:mb-5 text-2xl/7 md:text-4xl text-primary font-semibold text-center">Be Part of Something Bigger</h2>
          <Contact/>
        </section>
      </main>
  );
}
