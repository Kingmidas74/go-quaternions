# Quaternions

There is my simple implementation of Quaternions (<img src="https://latex.codecogs.com/svg.image?\mathbb{H}" alt="H"/>). Basic issues it resolves:
- Rotation some 3D point around axis
- Parallel displacement some 3D point
- Combination of parallel displacement with rotation
- And last but not least - rotation random 3D point around another 3D point

Axiom of quaternion algebra (described by Sir William Rowan Hamilton at 1843): 

<img src="https://latex.codecogs.com/svg.latex?i^2=j^2=k^2=ijk=-1" alt="basic quaternion axiom"/>

Also, I used to dual quaternions (based on dual numbers):

<img src="https://latex.codecogs.com/svg.latex?B=p+Iq" alt="B-quaternion"/>, где <img src="https://latex.codecogs.com/svg.latex?p,q%20\in%20\mathbb{H},%20I^2=0,%20I%20\neq%200" alt="B-quaternion"/>

All basic operations was hidden by Vec3 abstraction, so you can use it for your purpose.





