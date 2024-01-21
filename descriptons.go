package main

var introduction = `As your eyes flutter open, you find yourself lying on a bed of soft moss beneath a thick canopy of ancient trees. 
The air is crisp, and an eerie mist hangs in the air, cloaking the surroundings in a mysterious haze. 
The only sound you hear is the distant rustling of leaves and the occasional hoot of an unseen owl.
The grove is nestled between towering, gnarled trees with twisted branches that seem to whisper secrets to each other. 
Shafts of pale moonlight filter through the dense foliage, casting an ethereal glow on the damp ground. Strange, 
luminescent mushrooms dot the landscape, providing a dim but otherworldly illumination.
As you sit up, you realize you wear unfamiliar, tattered clothing, and your body feels strangely resilient. There's a lingering sense of disorientation, 
and you can't recall how you arrived here. The memories are elusive, slipping away like mist through your fingers.
Ahead, a narrow path winds its way through the ancient trees, disappearing into the fog. To the left, a babbling brook reflects the soft moonlight, 
and to the right, a mysterious archway stands, partially obscured by vines and foliage. Beyond the arch, shadows dance, suggesting the presence of something hidden.
The air is thick with anticipation, and a sense of both curiosity and trepidation washes over you. You stand at the precipice of the unknown, 
with the grove offering no clues about your past or the purpose of your presence here.

What would you like to do?
Follow the path: Venture deeper EAST into the grove, hoping to unravel the mysteries that shroud your memory.
Investigate the archway: Examine the enigmatic arch, curious about what lies beyond.
Explore the brook: Follow the gentle sounds of running water to the WEST and discover what secrets the brook may hold.`

var tile_32_short = ``
var tile_32_description = `Stepping through the moss-laden stones, you emerge from the emerald embrace of the twilight grove. 
Sunlight, dappled and warm, filters through the canopy overhead, painting the forest floor in a mosaic of light and shadow. 
The oppressive mist that shrouded the grove like a spectral veil thins and dissipates, revealing a world reborn in vibrant hues.  
Gone are the gnarled branches clawing at the sky, replaced by a cathedral of towering trees. 
Sunlight bathes their emerald leaves in liquid gold, casting long, dancing shadows across the verdant carpet below. 
The air, once thick with the cloying scent of damp earth, now hums with the sweet symphony of birdsong and the rustling whispers of leaves.
A carpet of wildflowers, a kaleidoscope of amethyst and saffron, stretches before you. Butterflies, their wings like stained glass windows, 
flit from bloom to bloom, their vibrant colors a stark contrast to the muted tones of the grove. A gentle breeze stirs the emerald sea, 
sending ripples of color dancing across the floral expanse.  In the distance, a crystalline stream meanders through the glade, 
its surface a mirror reflecting the azure sky above. Sunlight dances on the water, transforming it into a liquid tapestry of diamonds and sapphires. 
The gurgle of its current mingles with the birdsong, creating a melody that soothes the soul.  A path, worn smooth by the tread of countless unseen feet, 
beckons you deeper into the sun-dappled heart of the glade. The ancient stones that pave its way are etched with cryptic symbols, 
their meaning lost to the ages yet pulsing with an undeniable magic. The air here crackles with an unseen energy, 
a whispered promise of adventure that sets your heart racing with anticipation.  As you step onto the path, the sunlight warms your skin, 
chasing away the lingering chill of the grove. You take a deep breath, the fresh, sweet air filling your lungs, and a sense of wonder washes over you. 
This is no ordinary glade; it is a portal to a world of vibrant life, ancient secrets, and endless possibilities.  It's then, with the first tentative 
steps on this sunlit haven, that a flicker of recognition ignites within. A spark in the fog-choked memory that once held your past. 
Slowly, like an ancient scroll unfurling, fragmented glimpses emerge: a final breath, a searing flash, a peaceful surrender. The pieces click into place, 
forming a horrifying yet liberating truth - you are no longer bound by the mortal coil. This sunlit glade, this vibrant symphony of life, it's not just a paradise found, 
it's a realm beyond the veil, a land of the ever-after.  A strange calmness washes over you. Fear, perhaps, should gnaw at the edges of your being, 
yet an undeniable acceptance takes hold. The sun warms your skin, not with mere light, but with a gentle understanding. 
The rustling leaves whisper not secrets, but echoes of countless souls who have passed before. Perhaps this death is not an ending, but a homecoming, 
a transition into a new chapter of existence, whispered about in hushed tones throughout life's fleeting moments.  With newfound clarity, you stand tall, 
not in defiance of your fate, but in humble acceptance. This glade, bathed in eternal sunlight, may not be the world you knew, but it is yours to explore, 
to understand. The ancient path beckons, not with promises of earthly destinations, but with the whispers of untold mysteries, tales woven from the fabric of eternity itself. 
This is not a path you fear, but one you embrace, a first step into an existence unbound by the limitations of mortal clay.  The sun shines bright, 
not just on the glade, but on your newly awakened soul. You raise your head, a smile gracing your lips. This is not an end, but an endless beginning, 
a chance to dance in the sunlit glade of forever. It's time to walk the path, not with the trepidation of the lost, but with the curiosity of the explorer, 
the wonder of the awakened.`

var tile_23_short = `You find yourself back at the brook.`
var tile_23_description = `As you decide to explore the babbling brook to your left, you make your way through the damp, moss-covered ground, drawn by the soft sounds of running water. 
The moonlight filters through the dense canopy, casting a gentle glow on the scene. The brook appears enchanting at first glance, its crystal-clear waters reflecting the pale moonlight, 
creating a mesmerizing dance of ripples.
The brook is lined with smooth stones and surrounded by vibrant, luminescent undergrowth that seems to sway in harmony with the water's melody. Ethereal fireflies dance above the surface, 
adding to the magical ambiance. It's a picturesque sight that captivates your senses.
As you approach the edge, you notice the water's surface shimmers with an otherworldly glow. A soft, enchanting hum emanates from the brook, enticing you to step closer. However, 
as you extend your hand to touch the water, an invisible barrier halts your progress. The enchantment that gives the brook its captivating beauty also makes it impassable.
Despite the surface beauty, an unsettling feeling begins to stir within you. An eerie unease takes root in your mind, as if unseen eyes are watching your every move. 
The vibrant flora surrounding the brook, once enchanting, now seems to writhe with an unnatural energy. Shadows dance beneath the moonlit surface, 
hinting at a darkness that contrasts sharply with the initial allure.
It becomes clear that the brook, while enchanting, holds a duality—a dichotomy of beauty and an underlying darkness. The air is thick with an unsettling tension, 
and a choice presents itself:

Attempt to dispel the enchantment: EXAMINE ways to overcome the magical barrier and explore the mysterious depths beyond.
Return to the EAST: Leave the brook and consider other paths within the grove.`
var tile_23_examine = `As you focus your attention on the brook, you notice intricate patterns in the smooth stones that line its edges. 
The luminescent undergrowth seems to respond to the water's flow, casting ethereal reflections on the stones. 
The fireflies overhead create a celestial dance, adding to the enchanting spectacle.  The invisible barrier preventing your progress appears 
to be an integral part of the brook's magical nature. Your attempts to force your way past the enchantment reveal the strength 
and complexity of the magical weave that surrounds it. As you push against the barrier, a surge of resistance pushes back, almost like an unseen hand preventing you from proceeding.
However, as you persist in trying to force your way through, a growing sense of foreboding takes hold. The once-mesmerizing fireflies now seem to flicker erratically, 
creating an erratic pattern of light and shadow.  A whispering wind carries subtle, eerie voices that seem to caution against pushing further. 
The ancient trees' branches overhead seem to reach out, casting elongated, spectral shapes on the ground. 
The enchantment reveals a darker side, hinting at consequences that may befall those who defy its intended purpose.
A choice now confronts you:

Continue pressing on: Despite the ominous signs, you can persist in trying to force your way past the magical barrier and explore what lies beyond.
Decide to stop: Heed the warning of the brook's enchanted nature and choose to turn away, considering alternative paths within the grove.`


var tile_33_short = "You recognize this place as the spot where consciousness returned, marked by the presence of the imposing archway that looms with an air of intimidation."
var tile_33_examine = `As you approach the mysterious archway to your right, the scene unfolds before you. The arch stands tall and imposing, its structure composed of ancient, 
weathered stones that seem to have withstood the passage of countless years. Vines and foliage partially obscure the arch, adding to its enigmatic aura.
The stones that make up the archway bear intricate carvings and symbols, telling a silent tale of an age long past. Each rune and symbol appears deliberate, 
contributing to the overall mystique of the structure. As you run your fingers along the stones, you notice that one particular rune stands out—an oddity 
amidst the otherwise cohesive patterns.
This peculiar rune seems to be missing a piece, a strangely shaped stone that once completed its intricate design. The absence is subtle but noticeable, 
creating a sense of imbalance within the rune. The remaining stones around it depict a cohesive narrative, but this one rune hints at an unresolved mystery.
Upon closer inspection, you observe that the rune with the missing piece doesn't quite match the others. Its edges appear jagged, 
as if a part of it had been intentionally removed or lost over time. The rune seems to yearn for completion, and the empty space within it suggests that 
something uniquely shaped once filled that void.  The air around the archway carries a faint resonance, as if the missing piece held significance beyond mere aesthetics. 
Shadows dance on the stones as if whispering secrets, and the atmosphere is pregnant with untold stories. 
A choice presents itself:
Search for the missing piece: EXAMINE the surroundings to find the strangely shaped stone that completes the rune, unveiling the archway's secrets.
Continue exploring: Leave the archway for now and explore other areas of the grove, keeping the mysterious rune in mind.`

var tile_43_short = "Once again, you stand before the enigmatic stone structure."
var tile_43_description = `You decide to follow the narrow path that winds through the ancient trees, disappearing into the thick fog. 
Each step you take feels both uncertain and deliberate as the mist envelops you, reducing visibility to mere shadows and outlines.
The atmosphere becomes denser with each stride, and the air is filled with the sweet scent of moss and damp earth. 
Occasionally, you catch glimpses of ethereal shapes flitting through the fog – perhaps creatures of the grove or figments of your imagination.
As you venture deeper, the twisted branches overhead form a natural canopy, creating an almost tunnel-like effect. The path occasionally meanders, 
revealing patches of vibrant flowers and curious flora. The mysterious aura of the grove intensifies, and you sense an unseen presence watching your every move.
After some time, the fog begins to thin, and you find yourself in a small clearing. Here, the moonlight pierces through the remaining mist, 
illuminating a peculiar stone structure at the center. Symbols and runes adorn its surface, hinting at an ancient significance lost to time.
EXAMINE the stone structure: Investigate the mysterious symbols and runes, seeking clues about the grove's history.
Continue along the path: Press forward, determined to uncover the secrets that still linger within the fog.
Retrace your steps: Return to the heart of the grove, considering other paths that may unravel the mysteries shrouding your memory.`
var tile_43_examine = `You closely examine the stone structure, paying particular attention to the ancient symbols and runes etched into its surface. 
Running your fingers over the weathered stones, you notice a peculiar rune that stands out – it has an oddly shaped stone at its center, 
appearing deliberate, as if designed to hold something specific.  The stone structure exudes a faint energy, 
and you sense that uncovering the purpose of this peculiar stone might reveal hidden secrets.

Options:
Try and take the stone.
Continue exploring: Leave the stone structure for now and continue your journey through the grove, keeping the peculiar rune in mind.`


var stone_placed = `As you delicately place the peculiarly shaped stone into the rune's oddly shaped indent, you feel a subtle shift in the air. 
The stone fits into the void with an almost magnetic pull, settling seamlessly into place. As the stone finds its spot, a hushed vibration echoes through the grove, 
and you sense a connection resonating with the wild flora that surrounds you.  The ancient trees seem to stir, 
their gnarled branches shifting in response. The luminescent mushrooms emit a soft glow, synchronizing with the newfound energy.
A chorus of unseen creatures, perhaps hidden within the shadows, releases a symphony of chirps and rustles. 
It's as if the grove itself is awakening, responding to the adjustment you made to the stone structure.
A gentle tremor, like the heartbeat of the grove, emanates from the ground. It's a harmonious dance between nature and the mysterious 
forces that linger in the air. The once cloaked surroundings now reveal a heightened vitality, and you find yourself at the center 
of this enigmatic convergence.  With the stone in place, you notice subtle changes in the archway to the right. 
The vines and foliage part slightly, revealing a passage beyond. Shadows retreat, unveiling a path that leads deeper into the heart of the grove. The air, charged with newfound energy, beckons you to explore further.

What would you like to do?
Venture NORTH through the revealed passage: Follow the path beyond the archway, curious about the secrets it might unveil.
Explore the changed surroundings: Observe the altered grove, paying attention to any new details that may have emerged.
Head back to the brook or the path: Return to familiar locations, either the babbling brook to the left or the narrow path winding through the ancient trees.`

var help = `COMMANDS:  NORTH, SOUTH, EAST, WEST, EXAMINE, GET ITEM, PLACE, INVENTORY, HELP, QUIT`

var hellscape = `As you persist in attempting to force your way past the brook's enchantment, the air thickens with an unsettling energy.  
A haunting hum resonates, growing louder with each futile attempt. The smooth stones lining the brook's edge reveal hidden mouths emitting tortured whispers. 
The ancient trees' branches contort, casting skeletal shadows that reach out as if trying to warn you.
As the barrier gives way, you find yourself being drawn into the brook's depths. The water transforms into a turbulent whirlpool that pulls you into a realm of distorted 
visions—screaming souls, contorted faces, and indescribable horrors. The air is thick with despair, and the haunting hum becomes a cacophony of wails.
The once-beautiful brook has become a gateway to a dimension of eternal suffering. Tormented spirits claw at you, their spectral hands reaching out in silent agony. 
Your senses are overwhelmed by unbearable suffering. The consequences of defying the brook's enchantment are clear.
In this realm, time loses meaning, and the pain seems eternal. The enchantment, now a prison of despair, serves as a warning to those who defy the natural order. 
Your existence within this nightmarish dimension becomes a testament to the consequences of pushing past ancient magic's boundaries.
The choice to press on has led to an irreversible fate—a descent into an abyss of agony, with no escape in sight.  Your soul is now ensnared within this inescapable realm of torment.`