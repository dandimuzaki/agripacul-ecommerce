import { formatTime } from "@/lib/formatDate";
import { OrderStep } from "@/types/order";

const OrderStatus = ({ step }: {step: OrderStep}) => {
  return (step.done &&
    (<div className='flex gap-4'>
      <div className='flex flex-col items-center justify-center relative'>
        <div className={`${(step.key == 'created') ? 'bg-none' : 'bg-primary'} 'flex-1 w-1 h-full absolute top-[-50%]`}></div>
        <div className='h-3 w-3 rounded-full bg-primary absolute top-[50%] left-[50%] -translate-1/2'></div>
      </div>
      <div className='flex-1 grid gap-1 my-2'>
        <div className='font-bold'><span className="text-primary">{step.label}</span> at {formatTime(step.at)}</div>
        <p>{String(step.description)}</p>
      </div>
    </div>)
  );
};

export default OrderStatus;