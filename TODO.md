# TODO

## 1. Getting camera data

- [ ] Build a simple static test scene in Blender (a room, a few objects)
- [ ] Place 3–4 cameras with different positions, focal lengths, and resolutions
- [ ] Add marker points at known positions (and a few known distances between them)
- [ ] Render an image from each camera
- [ ] Export ground truth per camera: position, rotation, focal length, sensor size, resolution
- [ ] Export ground truth marker positions (3D) and their pixel locations in each image (2D)
- [ ] (Optional) Export per-pixel depth maps for later comparison
- [ ] Decide on a file format for scene/camera data that the main program can read

## 2. Using points to find camera positions

- [ ] Write the pinhole camera model: project a 3D point to a 2D pixel
- [ ] Verify projection against Blender's exported pixel locations
- [ ] Given known 3D markers + their 2D pixels, solve for a single camera's position and rotation
- [ ] Compare the solved pose against ground truth and measure the error
- [ ] Solve for the camera's focal length too (not just pose)
- [ ] Replace known 3D marker positions with only known *distances* between markers
- [ ] Solve all cameras relative to each other, using the known distances to set real-world scale
- [ ] Add noise to the 2D marker locations and see how error grows
- [ ] Add lens distortion and see how it breaks things

## 3. What can we do with the pixel data

- [ ] For a pixel in one camera, compute its ray into the scene
- [ ] Draw that ray's line (epipolar line) in another camera's image
- [ ] Triangulate a point seen in two cameras and compare to ground truth
- [ ] Try matching the same point across cameras automatically (no hand-labeled markers)
- [ ] Try a coarse voxel grid: keep voxels whose color agrees across cameras that see them
- [ ] Compare the result against Blender's depth maps / geometry
- [ ] Output something viewable (point cloud or voxels) to inspect the result

## Open questions

- [ ] Core 3D representation: voxels, point cloud, distance field, or several?
- [ ] How many cameras, and what placement, is "enough"?
- [ ] How much can geometry alone recover before appearance/priors are needed?
