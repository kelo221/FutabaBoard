<script lang="ts">
  let posX = 0;
  let posY = 0;

  let isDragging = false;
  let offsetX = 0;
  let offsetY = 0;

  function dragStart(event) {
    isDragging = true;
    offsetX = event.clientX - posX;
    offsetY = event.clientY - posY;
  }

  function dragStop() {
    isDragging = false;
  }

  function dragMove(event) {
    if (isDragging) {
      posX = event.clientX - offsetX;
      posY = event.clientY - offsetY;
    }
  }
</script>

<style>
    .draggable-box {
        position: absolute;
        width: 100px;
        height: 100px;
        background-color: #3498db;
        cursor: grab;
    }

    .draggable-box:active {
        cursor: grabbing;
    }
</style>

<div
  class="draggable-box"
  style="top: {posY}px; left: {posX}px;"
  on:mousedown={dragStart}
  on:mouseup={dragStop}
  on:mousemove={dragMove}
>
</div>