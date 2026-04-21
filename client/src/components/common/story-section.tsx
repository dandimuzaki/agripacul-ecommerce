import { CheckCircle } from "@mui/icons-material"

const StorySection = () => {
  return (
    <div className="grid md:grid-cols-2 gap-6 text-white">
      <div className="text-sm md:text-base flex flex-col items-center md:block">
        <div className="mb-4 text-white bg-white/20 px-3 py-1 rounded-full font-medium w-fit uppercase text-xs md:text-sm">Our Story</div>
          <h2 className="text-center md:text-left mb-3 md:mb-5 text-2xl md:text-4xl text-primary font-bold text-center md:text-left">Why We Grow</h2>
          <div className="text-left">
          <p>We started with 2020m² of land and a simple idea:</p>
          <p className="font-bold">Food should be fresh, accessible, and meaningful.</p>
          <br/>
          <p>What started as a small cultivation effort has grown into something more intentional. We are not only producing vegetables, but also building a closer connection between people and the food they consume.</p>
          <br/>
          <p><span className="text-lg md:text-xl text-primary uppercase font-bold">Agripacul is more than farming,</span> it's about</p>
          
          <ul>
            <li className="flex gap-2 items-center">
              <div className="text-primary text-lg md:text-xl"><CheckCircle/></div>
              Supporting local agriculture
            </li>
            <li className="flex gap-2 items-center">
              <div className="text-primary text-lg md:text-xl"><CheckCircle/></div>
              Making healthy food practical for students</li>
            <li className="flex gap-2 items-center">
              <div className="text-primary text-lg md:text-xl"><CheckCircle/></div>
              Bringing farming closer to everyday life</li>
          </ul>
          <br/>
          <p>From soil to salad, we stay connected to every step.</p>
          </div>
      </div>
      <div className="h-full md:relative grid gap-2 md:block">
        <div className="relative rounded-lg overflow-hidden h-full w-full md:absolute right-0">
          <img src="/story.png" alt="Fresh Vegetables by Agripacul" width={100} height={100} className="h-full w-full object-cover relative object-right" />
          {/* <div className="p-6 flex flex-col gap-1 justify-end absolute z-2 bottom-0 h-full w-full bg-[linear-gradient(rgba(0,0,0,0),rgba(0,0,0,0),rgba(0,0,0,0.6))]">
          </div> */}
        </div>
        {/* <div className="md:absolute md:top-8 md:bottom-8 md:left-0 grid grid-cols-2 md:grid-cols-1 md:grid-rows-2 gap-2 md:gap-4 z-3 md:w-42">
          <div className="text-black p-6 rounded bg-[rgb(202,241,229)] h-full w-full flex justify-center flex-col">
            <p className="text-2xl md:text-5xl font-medium">50+</p>
            <p className="text-sm md:text-base/5">products we provide</p>
          </div>
          <div className="p-6 rounded bg-primary text-white h-full w-full flex justify-center flex-col">
            <p className="text-2xl md:text-5xl font-meidum">100+</p>
            <p className="text-sm md:text-base/5">happy customers</p>
          </div>
        </div> */}
      </div>
    </div>
  )
}

export default StorySection
