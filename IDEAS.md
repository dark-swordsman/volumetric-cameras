This is a document just to explain my ideas about what this project is and its goals.

Primarily: The idea is to take any number of *any* camera and try to paint a volumetric space from it.

Why?

There are already tools to determine the lens/focal information of a camera. On top of that, combined points are often used track things, like mocap.

The idea: If we *know* the real length of an object, or distance between points in a room, and we identify those same points in a given camera view, and then take each pixel and know where it *should* be in space relative to other cameras. I know this may not be perfect, but I think there is enough here to explore and play with.

The best part is that we should be able to much more quickly test this using a game engine or Blender for quick iteration, and then we can try to apply it IRL. 